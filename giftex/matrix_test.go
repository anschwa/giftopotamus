package giftex

import (
	"fmt"
	"testing"
)

func Example_matrix() {
	m := newMatrix(4)
	fmt.Printf("Identity:\n%v\n\n", m)

	c := constraints{
		0: []Pid{1},
		1: []Pid{0, 2},
		2: []Pid{2, 3},
		3: []Pid{0, 2, 3},
	}

	m.AddConstraints(c)
	fmt.Printf("Constraints:\n%v\n\n", m)

	if m.CheckConstraints() {
		a := m.Assign()
		fmt.Printf("Assignment:\n%v\n\n", a)
		fmt.Printf("Valid? %v\n", verifyAssignment(a, c))
	} else {
		fmt.Println("No solution")
	}

	// Output:
	// Identity:
	// 1 0 0 0
	// 0 1 0 0
	// 0 0 1 0
	// 0 0 0 1
	//
	// Constraints:
	// 1 1 0 0
	// 1 1 1 0
	// 0 0 1 1
	// 1 0 1 1
	//
	// Assignment:
	// 0 -> 2
	// 1 -> 3
	// 2 -> 0
	// 3 -> 1
	//
	// Valid? true
}

func cmpMatrix(t *testing.T, want, got matrix) {
	size := len(want)
	if a, b := len(got), len(want); a != b {
		t.Errorf("wrong number of rows for size %d:\nwant: %d; got: %d", size, a, b)
	}

	for i := range got {
		if a, b := len(got[i]), len(want[i]); a != b {
			t.Errorf("wrong number of columns for size %d:\nwant: %d; got: %d", size, a, b)
		}

		for j := range got[i] {
			if want[i][j] != got[i][j] {
				t.Errorf("wrong matrix for size %d:\nwant:\n%v\n\ngot:\n%v", size, want, got)
			}
		}
	}
}

func TestNewMatrix(t *testing.T) {
	// NewMatrix(N) should return the identity matrix of dimension N

	tests := []struct {
		size int
		want matrix
	}{
		{size: 0, want: [][]int{}},
		{size: 1, want: [][]int{{1}}},
		{size: 2, want: [][]int{{1, 0}, {0, 1}}},
		{size: 3, want: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
	}

	for _, tt := range tests {
		cmpMatrix(t, tt.want, newMatrix(tt.size))
	}
}

func TestMatrix_String(t *testing.T) {
	tests := []struct {
		size int
		want string
	}{
		{0, ""},
		{1, "1"},
		{2, "1 0\n0 1"},
		{3, "1 0 0\n0 1 0\n0 0 1"},
	}

	for _, tt := range tests {
		got := newMatrix(tt.size).String()
		if tt.want != got {
			t.Errorf("wrong output for size %d:\nwant: %#v; got: %#v", tt.size, tt.want, got)
		}
	}
}

func TestMatrix_AddConstraints(t *testing.T) {
	tests := []struct {
		size int
		c    constraints
		want matrix
	}{
		{
			size: 3,
			c:    map[Pid][]Pid{0: {0, 1, 2}, 1: {}, 2: {0}},
			want: [][]int{{1, 1, 1}, {0, 1, 0}, {1, 0, 1}},
		},
	}

	for _, tt := range tests {
		m := newMatrix(tt.size)
		m.AddConstraints(tt.c)

		cmpMatrix(t, tt.want, m)
	}
}

func TestMatrix_CheckConstraints(t *testing.T) {
	tests := []struct {
		m    matrix
		want bool
	}{
		{m: [][]int{{1, 1}, {1, 1}}, want: false},
		{m: [][]int{{0, 0}, {0, 0}}, want: true},
		{m: [][]int{{1, 1}, {0, 1}}, want: false},
		{m: [][]int{{1, 0}, {0, 1}}, want: true},
		{m: [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}, want: false},
		{m: newMatrix(3), want: true},
	}

	for i, tt := range tests {
		if got := tt.m.CheckConstraints(); got != tt.want {
			t.Errorf("want: %v; got: %v for matrix %d\n%v", tt.want, got, i, tt.m)
		}
	}
}

func TestAssignment_String(t *testing.T) {
	a := Assignment{
		2: 3,
		1: 2,
		0: 1,
	}

	want := "0 -> 1\n1 -> 2\n2 -> 3"
	if got := a.String(); want != got {
		t.Errorf("want:\n%#v\ngot:\n%#v", want, got)
	}
}

func TestMatrix_Assign(t *testing.T) {
	size := 3
	m := newMatrix(size)

	// Pick constraints that force only one choice for each assignment
	// since all we care about here is that the constraints are met.
	c := constraints{
		0: []Pid{2},
		1: []Pid{0},
		2: []Pid{1},
	}

	m.AddConstraints(c)
	a := m.Assign()

	// Check that we have one assignment for each participant
	if got := len(a); size != got {
		t.Errorf("wrong number of assignments: want: %d; got: %d", size, got)
	}

	// Check that assignments meet all constraints
	for p, exclusions := range c {
		for _, e := range exclusions {
			if a[p] == e {
				t.Errorf("failed constraint: %v cannot be assigned to %v", p, e)
			}
		}
	}
}

func TestMatrix_Verify(t *testing.T) {
	tests := []struct {
		a    Assignment
		c    constraints
		want bool
	}{
		{
			a:    map[Pid]Pid{0: 42},
			c:    map[Pid][]Pid{0: {42}},
			want: false,
		},
		{
			a:    map[Pid]Pid{0: 42},
			c:    map[Pid][]Pid{},
			want: true,
		},
	}

	for i, tt := range tests {
		if got := verifyAssignment(tt.a, tt.c); tt.want != got {
			t.Errorf("assignment %d is invalid: want: %v; got: %v", i, tt.want, got)
		}
	}
}
