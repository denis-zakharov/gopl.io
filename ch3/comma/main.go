// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
		fmt.Printf("  %s\n", comma2(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	commas := n/3 - 1
	rem := n % 3
	var buf *bytes.Buffer = bytes.NewBuffer(make([]byte, commas))
	if rem > 0 {
		buf.WriteString(s[:rem])
		buf.WriteString(",")
	}
	for i := rem; i < n-3; {
		next := i + 3
		buf.WriteString(s[i:next])
		buf.WriteString(",")
		i += 3
	}
	buf.WriteString(s[n-3 : n])

	return buf.String()
}
