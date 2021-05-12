// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

/* 32 or 64: platform dependnent
Switch on all bits, right shift to 63 gives:
0b1 on 64-bit plaforms or
0b0 on 32-bit platforms.
32<<0 does nothing.
32<<1 is the same as 32*2.
*/
const uintSize = 32 << (^uint(0) >> 63)

func PopCount(x uint) int {
	res := 0
	for x != 0 {
		x &= x - 1
		res += 1
	}
	return res
}

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Len returns a number of elements is in the set.
func (s *IntSet) Len() int {
	res := 0
	for _, tword := range s.words {
		if tword != 0 {
			res += PopCount(tword)
		}
	}
	return res
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintSize, uint(x%uintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds variate number of integers to the set.
func (s *IntSet) AddAll(xs ...int) {
	for x := range xs {
		s.Add(x)
	}
}

// Remove removes x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// Removes all elements from the set.
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	var newWords = make([]uint, len(s.words))
	copy(newWords, s.words)
	return &IntSet{newWords}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	if len(t.words) < len(s.words) {
		s.words = s.words[:len(t.words)]
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, sword := range s.words {
		if i < len(t.words) {
			s.words[i] = sword &^ t.words[i]
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []int {
	iterator := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				iterator = append(iterator, uintSize*i+j)
			}
		}
	}
	return iterator
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
