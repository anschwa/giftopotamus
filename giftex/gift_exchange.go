package giftex

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
)

type GiftExchange struct {
	numParticipants int
	matrix          matrix
	participants    ParticipantMap
	constraints     constraints
	Assignment      Assignment
}

var (
	ErrInvalidCSV          = errors.New("Error: csv must include headers and at least one entry")
	ErrNoSolution          = errors.New("No Solution: an assignment is not possible for this gift exchange")
	ErrParticipantNotFound = errors.New("Error: Participant not found in GiftExchangeDB")
)

type GiftExchangeOptions struct {
	// MaxPrevious is how long to wait until you can pair with someone you had before
	MaxPrevious int
}

func NewGiftExchange(pm ParticipantMap, opts *GiftExchangeOptions) (*GiftExchange, error) {
	n := len(pm)
	m := newMatrix(n)

	var maxPrev int
	if opts != nil {
		maxPrev = opts.MaxPrevious
	}

	c := make(constraints, n)
	for id, x := range pm {
		// Previous assignments older than maxPrev are allowed
		if n := len(x.Previous); maxPrev > 0 && n > maxPrev {
			c[id] = append(x.Restrictions, x.Previous[n-maxPrev:]...)
			continue
		}

		c[id] = append(x.Restrictions, x.Previous...)
	}

	ge := &GiftExchange{
		numParticipants: n,
		matrix:          m,
		participants:    pm,
		constraints:     c,
	}

	if ok := m.AddConstraints(c); !ok {
		return ge, ErrNoSolution
	}

	// It's possible to "choose wrong" when making an assignment that
	// will prevent a complete pairing to be made. Let's give
	// ourselves as many tries as we have participants to choose correctly.
	var a Assignment
	var valid bool
	for i := 0; i < n; i++ {
		a = m.Assign()
		if verifyAssignment(a, c) {
			valid = true
			break
		}
	}

	if valid {
		ge.Assignment = a
	} else {
		return nil, ErrNoSolution
	}

	return ge, nil
}

type ParticipantMap = map[Pid]Participant
type Participant struct {
	ID           Pid
	Name         string
	Email, SMS   string
	Restrictions []Pid
	Previous     []Pid
}

type GiftExchangeDB struct {
	cols    map[string]int
	headers []string
	records [][]string

	Participants ParticipantMap
	index        map[Pid]int
}

func (db *GiftExchangeDB) GetParticipant(id Pid) (Participant, error) {
	p, ok := db.Participants[id]
	if !ok {
		return Participant{}, ErrParticipantNotFound
	}

	return p, nil
}

func (db *GiftExchangeDB) loadRecords() {
	numRecords := len(db.records)
	nameMap := make(map[string]Pid, numRecords)

	var pID Pid
	for i := 0; i < numRecords; i++ {
		row := db.records[i]

		// Skip non-participants
		if participating := trimLower(row[db.cols["participating"]]) == "yes"; !participating {
			continue
		}

		p := Participant{
			ID:    pID,
			Name:  trim(row[db.cols["name"]]),
			Email: trimLower(row[db.cols["email"]]),
			SMS:   onlyDigits(row[db.cols["sms"]]),
		}

		db.Participants[p.ID] = p
		nameMap[trimLower(p.Name)] = p.ID
		db.index[pID] = i
		pID++ // Increment pID last so we start with 0
	}

	// Second pass to fill out constraints
	for pID, p := range db.Participants {
		getIDs := func(entry string) []Pid {
			names := strings.Split(trim(entry), ",")
			ids := make([]Pid, 0, len(names))

			for _, n := range names {
				if n == "" {
					continue
				}

				// Ignore names that are not real participants
				if id, ok := nameMap[trimLower(n)]; ok {
					ids = append(ids, id)
				}
			}

			return ids
		}

		idx := db.index[pID]
		p.Restrictions = getIDs(db.records[idx][db.cols["restrictions"]])
		p.Previous = getIDs(db.records[idx][db.cols["previous"]])
		db.Participants[pID] = p
	}
}

// ReadCSV processes a CSV representation of prior gift exchange data
// and preserves the original records for writing out as a new CSV later.
//
// The following columns are required: name, email, restrictions, previous, participating, has
func ReadCSV(r io.Reader) (*GiftExchangeDB, error) {
	csvReader := csv.NewReader(r)
	csvReader.FieldsPerRecord = -1 // Allow empty columns
	csvReader.Comma = ','

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error reading csv: %w", err)
	}

	if len(records) < 2 {
		return nil, ErrInvalidCSV
	}

	// The first record contains the column headers
	cols := make(map[string]int, len(records[0]))
	for i, v := range records[0] {
		cols[trimLower(v)] = i
	}

	maxSize := len(records) - 1 // Skip header row

	db := &GiftExchangeDB{
		cols:         cols,
		headers:      records[0],
		records:      records[1:],
		Participants: make(ParticipantMap, maxSize),
		index:        make(map[Pid]int, maxSize),
	}

	db.loadRecords()
	return db, nil
}

func ReadCSVFromFile(path string) (*GiftExchangeDB, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading csv: %w", err)
	}
	defer f.Close()

	return ReadCSV(f)
}

// WriteCSV produces a new CSV with the results of a completed gift
// exchange while also preserving the original CSV's data. The rows
// are sorted by name.
func (db *GiftExchangeDB) WriteCSV(w io.Writer, results Assignment) error {
	b := csv.NewWriter(w)

	// Clear out "has" column from previous run
	for i, row := range db.records {
		row[db.cols["has"]] = ""
		db.records[i] = row
	}

	// Update records
	for pID, has := range results {
		idx := db.index[pID]
		row := db.records[idx]

		hasName := db.Participants[has].Name
		row[db.cols["has"]] = hasName

		// Split string by ',' but ignore empty items
		splitNames := func(s string) []string {
			return strings.FieldsFunc(s, func(c rune) bool {
				return c == ','
			})
		}

		prev := row[db.cols["previous"]]
		prev = strings.Join(append(splitNames(prev), hasName), ",")
		row[db.cols["previous"]] = prev

		db.records[idx] = row
	}

	// Sort records by Name
	sort.Slice(db.records, func(i, j int) bool {
		a, b := db.records[i], db.records[j]
		return a[db.cols["name"]] < b[db.cols["name"]]
	})

	// Write the column headers
	b.Write(db.headers)

	// Write updated records
	for _, row := range db.records {
		b.Write(row)
	}

	b.Flush()
	return b.Error()
}

func trimLower(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func onlyDigits(s string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(s, "")
}

func (ge *GiftExchange) String() string {
	// Build a slice of participants sorted by name
	type result struct {
		id   Pid
		name string
		has  string
	}

	sorted := make([]result, 0, ge.numParticipants)
	for aID, bID := range ge.Assignment {
		a, b := ge.participants[aID], ge.participants[bID]
		sorted = append(sorted, result{
			id:   aID,
			name: a.Name,
			has:  b.Name,
		})
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		return a.name < b.name
	})

	// Print results of the assignment
	var b strings.Builder
	w := tabwriter.NewWriter(&b, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, "Results:")
	for _, p := range sorted {
		fmt.Fprintf(w, "| %s\t | %s\t |\n", p.name, p.has)
	}

	w.Flush()
	return b.String()
}
