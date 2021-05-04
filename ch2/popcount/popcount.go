// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	// one right bit shift corresponds to an integer division by 2
	for i := range pc {
		// higher-order bits plus the last one
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 returns the number of set bits of x (PopCount using a loop).
func PopCount2(x uint64) int {
	byte_chunks := 8 // in uint64
	res := 0
	// Calculate a 1-bit count in each byte-chunk (8 byte chunks in uint64).
	for i := 0; i < byte_chunks; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

// PopCount3 returns the number of set bits of x by using one right-shift
// over all positions (a naive methods without cashing).
func PopCount3(x uint64) int {
	res := 0
	bits := 64
	for i := 0; i < bits; i++ {
		res += int((x >> i) & 1)
	}
	return res
}

// PopCount4 returns the number of set bits of x by using the fact:
// x&(x-1) switch off rightest non-zero bit.
func PopCount4(x uint64) int {
	res := 0
	for x != 0 {
		x &= x - 1
		res += 1
	}
	return res
}

//!-
