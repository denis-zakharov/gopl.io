// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 9}, 2},
		{[]int{0, 10, 13, 145}, 4},
	}

	for _, test := range tests {
		var x IntSet
		for _, i := range test.input {
			x.Add(i)
		}
		if got := x.Len(); got != test.want {
			t.Errorf("%v.Len() == %d, wanted %d", &x, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		input  []int
		remove int
		want   string
	}{
		{[]int{0, 10, 13, 145}, 4, "{0 10 13 145}"},
		{[]int{0, 10, 13, 145}, 13, "{0 10 145}"},
		{[]int{0, 10, 145}, 145, "{0 10}"},
		{[]int{0, 10}, 0, "{10}"},
		{[]int{10}, 0, "{10}"},
		{[]int{10}, 10, "{}"},
		{[]int{}, 10, "{}"},
	}

	for _, test := range tests {
		var x IntSet
		for _, i := range test.input {
			x.Add(i)
		}
		x.Remove(test.remove)
		if got := x.String(); got != test.want {
			t.Errorf("%v.Remove(%d) == %s, wanted %s", &x, test.remove, got, test.want)
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		input []int
	}{
		{[]int{0, 10, 13, 145}},
		{[]int{0, 10, 145}},
		{[]int{0, 10}},
		{[]int{10}},
		{[]int{}},
	}
	var want = "{}"

	for _, test := range tests {
		var x IntSet
		for _, i := range test.input {
			x.Add(i)
		}
		x.Clear()
		if got := x.String(); got != want {
			t.Errorf("%v.Clear() == %s, wanted %s", &x, got, want)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		input []int
		want  string
	}{
		{[]int{0, 10, 13, 145}, "{0 10 13 145}"},
		{[]int{0, 10, 145}, "{0 10 145}"},
		{[]int{0, 10}, "{0 10}"},
		{[]int{10}, "{10}"},
		{[]int{0}, "{0}"},
		{[]int{}, "{}"},
	}

	for _, test := range tests {
		var x IntSet
		for _, i := range test.input {
			x.Add(i)
		}
		yptr := x.Copy()
		if got := yptr.String(); got != test.want {
			t.Errorf("%v.Copy() == %s, wanted %s", &x, got, test.want)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		self  []int
		other []int
		want  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{}"},
		{[]int{1}, []int{}, "{}"},
		{[]int{1, 9, 42, 144}, []int{9, 144}, "{9 144}"},
		{[]int{1, 9, 42, 144}, []int{0, 4, 9, 42, 100}, "{9 42}"},
	}

	for _, test := range tests {
		var x, y IntSet
		for _, v := range test.self {
			x.Add(v)
		}
		for _, v := range test.other {
			y.Add(v)
		}
		x.IntersectWith(&y)
		if got := x.String(); got != test.want {
			t.Errorf("%v.IntersectWith(%v) == %s, wanted %s", &x, &y, got, test.want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		self  []int
		other []int
		want  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{}"},
		{[]int{1}, []int{}, "{1}"},
		{[]int{1, 9, 42, 144}, []int{9, 144}, "{1 42}"},
		{[]int{1, 9, 42, 144}, []int{0, 4, 9, 42, 100}, "{1 144}"},
	}

	for _, test := range tests {
		var x, y IntSet
		for _, v := range test.self {
			x.Add(v)
		}
		for _, v := range test.other {
			y.Add(v)
		}
		x.DifferenceWith(&y)
		if got := x.String(); got != test.want {
			t.Errorf("%v.DifferenceWith(%v) == %s, wanted %s", &x, &y, got, test.want)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		self  []int
		other []int
		want  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{1}"},
		{[]int{1}, []int{}, "{1}"},
		{[]int{1, 9, 42, 144}, []int{9, 144}, "{1 42}"},
		{[]int{1, 9, 42, 144}, []int{0, 4, 9, 42, 100}, "{0 1 4 100 144}"},
	}

	for _, test := range tests {
		var x, y IntSet
		for _, v := range test.self {
			x.Add(v)
		}
		for _, v := range test.other {
			y.Add(v)
		}
		x.SymmetricDifference(&y)
		if got := x.String(); got != test.want {
			t.Errorf("%v.SymmetricDifference(%v) == %s, wanted %s", &x, &y, got, test.want)
		}
	}
}

func TestElems(t *testing.T) {
	var tests = []struct {
		input []int
		want  string
	}{
		{[]int{}, "[]"},
		{[]int{1}, "[1]"},
		{[]int{1, 999}, "[1 999]"},
		{[]int{1, 9, 42, 144}, "[1 9 42 144]"},
	}

	for _, test := range tests {
		var x IntSet
		for _, v := range test.input {
			x.Add(v)
		}
		if got := fmt.Sprintf("%v", x.Elems()); got != test.want {
			t.Errorf("%v.Elems() == %s, wanted %s", &x, got, test.want)
		}
	}
}
