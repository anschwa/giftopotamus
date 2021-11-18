package giftex

import (
	"fmt"
	"strings"
	"testing"
)

func Example() {
	db, err := ReadCSVFromFile("testdata/small.csv")
	if err != nil {
		panic(err)
	}

	opts := &GiftExchangeOptions{MaxPrevious: 2}
	ge, err := NewGiftExchange(db.Participants, opts)
	if err != nil {
		panic(err)
	}

	fmt.Println(ge)
	// Output:
	// Results:
	// | bar  | foo  |
	// | baz  | bar  |
	// | foo  | baz  |
}

func Example_noSolution() {
	// Create an impossible assignment
	m := ParticipantMap{
		0: {ID: 0, Name: "foo", Restrictions: []Pid{1}},
		1: {ID: 1, Name: "bar", Restrictions: []Pid{0}},
	}

	_, err := NewGiftExchange(m, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// No Solution: an assignment is not possible for this gift exchange
}

func TestNewGiftExchange(t *testing.T) {
	pm := ParticipantMap{
		0: Participant{ID: 0, Previous: []Pid{1, 2}},
		1: Participant{ID: 1},
		2: Participant{ID: 2},
	}

	t.Run("Allow previous assignments older than MaxPrevious", func(t *testing.T) {
		// An assignment is only possible if previous participants are valid according to opts
		opts := &GiftExchangeOptions{MaxPrevious: 1}
		if _, err := NewGiftExchange(pm, opts); err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("Exclude all previous assignments with opts are nil", func(t *testing.T) {
		if _, err := NewGiftExchange(pm, nil); err == nil {
			t.Error("Assignment should fail without setting MaxPrevious")
			return
		}
	})
}

func TestReadCSVFromFile(t *testing.T) {
	wantDB := &GiftExchangeDB{
		cols: map[string]int{
			"name":          0,
			"email":         1,
			"sms":           2,
			"restrictions":  3,
			"previous":      4,
			"participating": 5,
			"has":           6,
		},
		headers: []string{"name", "email", "sms", "restrictions", "previous", "participating", "has"},
		records: [][]string{
			{"foo", "foo@example.com", "555 111 1111", "foo", "", "yes", ""},
			{"bar", "bar@example.com", "(555) 222-2222", "", "baz", "yes", ""},
			{"baz", "baz@example.com", "555.333.3333", "", "", "yes", ""},
			{"quux", "quux@example.com", "5554444444", "foo, bar, baz", "", "no", ""},
		},
		Participants: map[Pid]Participant{
			0: {ID: 0, Name: "foo", Email: "foo@example.com", SMS: "5551111111", Restrictions: []Pid{0}},
			1: {ID: 1, Name: "bar", Email: "bar@example.com", SMS: "5552222222", Previous: []Pid{2}},
			2: {ID: 2, Name: "baz", Email: "baz@example.com", SMS: "5553333333", Restrictions: []Pid{}},
		},
		index: map[Pid]int{0: 0, 1: 1, 2: 2},
	}

	db, err := ReadCSVFromFile("testdata/small.csv")
	if err != nil {
		t.Error(err)
	}

	if want, got := fmt.Sprintf("%v", wantDB), fmt.Sprintf("%v", db); want != got {
		t.Errorf("DBs don't match:\nwant: %v\n got: %v", want, got)
	}
}

func TestWriteCSV(t *testing.T) {
	db, err := ReadCSVFromFile("testdata/small.csv")
	if err != nil {
		t.Error(err)
		return
	}

	var b strings.Builder
	results := Assignment{0: 2, 1: 0, 2: 1}
	if err := db.WriteCSV(&b, results); err != nil {
		t.Error(err)
		return
	}

	want := `name,email,sms,restrictions,previous,participating,has
bar,bar@example.com,(555) 222-2222,,"baz,foo",yes,foo
baz,baz@example.com,555.333.3333,,bar,yes,bar
foo,foo@example.com,555 111 1111,foo,baz,yes,baz
quux,quux@example.com,5554444444,"foo, bar, baz",,no,
`

	if got := b.String(); want != got {
		t.Errorf("Wrong output:\nwant:\n%v\n\ngot:\n%v", want, got)
	}
}
