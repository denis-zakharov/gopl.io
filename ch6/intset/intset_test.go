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
	var x IntSet
	if x.Len() != 0 {
		t.Error("Empty")
	}
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	if x.Len() != 4 {
		t.Error("4 elements")
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	if x.String() != "{}" {
		t.Error("Empty")
	}
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	if x.String() != "{1 9 42 144}" {
		t.Error("Added")
	}

	x.Remove(240)
	if x.String() != "{1 9 42 144}" {
		t.Error("Remove non-existing")
	}
	x.Remove(42)
	if x.String() != "{1 9 144}" {
		t.Error("Remove existing")
	}
}

func TestClear(t *testing.T) {
	var x IntSet

	x.Clear()
	if x.Len() != 0 {
		t.Error("Clear empty")
	}

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	x.Clear()
	if x.Len() != 0 {
		t.Error("Clear non-empty")
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	var y *IntSet

	y = x.Copy()
	if len(y.words) != 0 && y.Len() != 0 {
		t.Error("Copy empty")
	}

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	// fmt.Println(&x)

	y = x.Copy()
	if y.String() != "{1 9 42 144}" {
		t.Errorf("Copy non-empty: %v", &y)
	}
}

func TestIntersectWith(t *testing.T) {
	var x IntSet
	var z IntSet
	z.Add(1)
	x.IntersectWith(&z)
	if x.Len() != 0 {
		t.Errorf("empty intersectWith non-empty")
	}

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	var y IntSet

	x.IntersectWith(&y)
	if x.Len() != 0 {
		t.Errorf("Intersect with an empty set")
	}

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	y.Add(1)
	y.Add(2)
	x.IntersectWith(&y)

	if x.String() != "{1}" {
		t.Errorf("one")
	}

	x.Add(144)
	x.Add(9)
	x.Add(42)
	x.Add(2)
	x.IntersectWith(&y)

	if x.String() != "{1 2}" {
		t.Errorf("two")
	}
}

func TestDifferenceWith(t *testing.T) {
	var x IntSet
	var z IntSet
	z.Add(1)
	x.DifferenceWith(&z)
	if x.Len() != 0 {
		t.Errorf("empty differenceWith non-empty")
	}

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	var y IntSet

	x.DifferenceWith(&y)
	if x.Len() != 4 {
		t.Errorf("Diff with an empty set")
	}

	y.Add(1)
	y.Add(2)
	x.DifferenceWith(&y)

	if x.String() != "{9 42 144}" {
		t.Error("{1 9 42 144} \\ {1 2}")
	}

	y.Add(144)
	y.Add(9)
	y.Add(42)
	x.DifferenceWith(&y)

	if x.String() != "{}" {
		t.Error("{9 42 144} \\ {1 2 9 42 144}")
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x IntSet
	var z IntSet
	z.Add(1)
	x.SymmetricDifference(&z)
	if x.String() != "{1}" {
		t.Errorf("empty symmetric diff {1}")
	}

	x.Add(144)

	var y IntSet

	x.SymmetricDifference(&y)
	if x.String() != "{1 144}" {
		t.Errorf("{1 144} with an empty set")
	}

	y.Add(1)
	y.Add(2)
	x.SymmetricDifference(&y)

	if x.String() != "{2 144}" {
		t.Error("{1 144} symm diff {1 2}")
	}

	y.Add(144)
	y.Add(9)
	y.Add(42)
	x.SymmetricDifference(&y)

	if x.String() != "{1 9 42}" {
		t.Error("{2 144} \\ {1 2 9 42 144}")
	}
}
