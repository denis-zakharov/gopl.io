package main

import (
	"fmt"
	"os"
)

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, c := range s1 {
		m[c]++
	}
	for _, c := range s2 {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: main word1 word2")
		os.Exit(1)
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Printf("%q and %q are anagrams? %v", s1, s2, isAnagram(s1, s2))
}
