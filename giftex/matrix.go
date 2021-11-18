package giftex

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

var random *rand.Rand

func init() {
	// Seed a private PRNG for use in this package only
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// A gift exchange can be modeled with the identify matrix I_n, where
// n is the number of participants in the exchange. This is because
// the minimum constraint is that nobody should get paired with
// themselves. So we can always start off with an identity matrix and
// fill in any additional constraints after that.
//
// In our matrix, a zero means an assignment is available. A one means
// there is a constraint preventing assignment to that position.
type matrix [][]int

// newMatrix will create a new identity matrix of dimension n.
func newMatrix(n int) matrix {
	m := make(matrix, n, n)
	for i := range m {
		m[i] = make([]int, n, n)
		for j := range m[i] {
			if i == j {
				m[i][j] = 1
			} else {
				m[i][j] = 0
			}

		}
	}

	return m
}

func (m matrix) String() string {
	var b strings.Builder
	for i := range m {
		for j := range m[i] {
			fmt.Fprintf(&b, "%d", m[i][j])
			if j < len(m[i])-1 {
				b.WriteString(" ")
			}
		}

		if i < len(m)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

type Pid int // The index of a given row represents a participant in the gift exchange
type constraints map[Pid][]Pid

// AddConstraints fills out matrix m with the given constraints and
// returns true if an assignment is possible. AddConstraints will
// panic if the constraints c contain entries outside the bounds of matrix m.
func (m matrix) AddConstraints(c constraints) bool {
	for i := range m {
		exclusions, ok := c[Pid(i)]
		if !ok {
			continue
		}

		for _, p := range exclusions {
			m[i][p] = 1
		}
	}

	return m.CheckConstraints()
}

// CheckConstraints determines whether an assignment exists that can satisfy all
// constraints. It borrows a few steps from the "Hungarian Algorithm",
// but all we care about here is if an assignment is possible for the
// given matrix. We don't have any weights/costs so there is no need
// to reduce or find an "optimum" assignment.
func (m matrix) CheckConstraints() bool {
	// Let N be the dimensions of the square matrix M.
	//
	// To "star" a zero or "cover" a row or column simply means to
	// keep track of information in some data-structure, such as
	// adding an index to a map.
	//
	// (1) Find a zero in the matrix. If there are no starred zeros in
	// the same row or column, star it. Repeat for each element in the matrix.
	//
	// (2) Cover each column containing a starred zero. If N columns
	// are covered then we know an assignment is possible, so we're DONE.
	//
	// (3) Find an un-starred zero in an uncovered column and star it.
	// If there are no other starred zeros in that row, an assignment
	// exists, so we're DONE. Otherwise, cover this row and uncover
	// the column containing the starred zero. Repeat until every zero
	// has been checked.
	//
	// Step (4) If there still exists un-starred zeros that belong to
	// covered rows or columns, then an assignment isn't possible and we're DONE.

	type point struct{ row, col int }

	size := len(m)
	stars := make(map[point]struct{}, size*size)
	starredRows := make(map[int]bool)
	starredCols := make(map[int]bool)
	coveredRows := make(map[int]bool, size)
	coveredCols := make(map[int]bool, size)

	starInRowOrCol := func(p point) bool {
		return starredRows[p.row] || starredCols[p.col]
	}

	star := func(p point) {
		stars[p] = struct{}{}
		starredRows[p.row] = true
		starredCols[p.col] = true
	}

	// Steps (1) and (2)
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 1 {
				continue
			}

			p := point{i, j}
			if !starInRowOrCol(p) {
				star(p)
				coveredCols[p.col] = true
			}
		}
	}

	// If there are no zeros at all, then we can't make an assignment
	if len(stars) == 0 {
		return false
	}

	// Exit early if we already covered all the columns
	if len(coveredCols) == size {
		return true
	}

	// Step (3)
	for i := range m {
		for j := range m[i] {
			p := point{i, j}
			if m[i][j] == 1 || coveredRows[i] || coveredCols[j] {
				continue
			}

			if !starredRows[i] {
				return true
			}

			// Uncover column of stared zero in this row
			for c := range m[i] {
				if _, starred := stars[point{i, c}]; coveredCols[c] && starred {
					coveredCols[c] = false
					break
				}
			}

			star(p)
			coveredRows[p.row] = true
		}
	}

	// Step (4)
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 0 && (coveredRows[i] || coveredCols[j]) {
				return false
			}
		}
	}

	return true
}

type Assignment map[Pid]Pid

func (a Assignment) String() string {
	// Sort assignments before printing
	sortedKeys := make([]Pid, 0, len(a))
	for p := range a {
		sortedKeys = append(sortedKeys, p)
	}
	sort.SliceStable(sortedKeys, func(i, j int) bool {
		return sortedKeys[i] < sortedKeys[j]
	})

	var b strings.Builder
	for i, k := range sortedKeys {
		fmt.Fprintf(&b, "%d -> %d", k, a[k])

		if i < len(sortedKeys)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func (m matrix) Assign() Assignment {
	participants := len(m)
	finalAssignment := make(Assignment, participants)
	alreadyAssigned := make(map[Pid]struct{}, participants)

	// The optimal way to make assignments is to sort our matrix by
	// rows containing the least to most available choices (i.e., "zeros").
	zeros := make(map[int]int, participants)
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 0 {
				zeros[i]++
			}
		}
	}

	type sortable struct{ index, count int }
	sortByZeros := make([]sortable, 0, participants)
	for i, count := range zeros {
		sortByZeros = append(sortByZeros, sortable{i, count})
	}

	sort.SliceStable(sortByZeros, func(i, j int) bool {
		return sortByZeros[i].count < sortByZeros[j].count
	})

	// Iterate over ordered indices to get optimal assignment
	for _, row := range sortByZeros {
		// Find available participants for assignment
		choices := make([]Pid, 0, participants)
		for j := range m[row.index] {
			if m[row.index][j] == 0 {
				choices = append(choices, Pid(j))
			}
		}

		// See if these participants are still available
		available := make([]Pid, 0, len(choices))
		for _, p := range choices {
			if _, taken := alreadyAssigned[p]; !taken {
				available = append(available, p)
			}
		}

		// Make a random choice if there is more than one option
		if lenAvail := len(available); lenAvail > 0 {
			randIdx := random.Intn(lenAvail)
			choice := Pid(available[randIdx])
			finalAssignment[Pid(row.index)] = choice
			alreadyAssigned[choice] = struct{}{}
		}
	}

	return finalAssignment
}

func verifyAssignment(a Assignment, c constraints) bool {
	// We should always have at least as many assignments as we have constraints.
	if len(a) < len(c) {
		return false
	}

	for p, exclusions := range c {
		for _, e := range exclusions {
			if a[p] == e {
				return false
			}
		}
	}

	return true
}
