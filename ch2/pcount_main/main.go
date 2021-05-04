package main

import (
	"fmt"

	"gopl.io/denis-zakharov/ch2/popcount"
)

func main() {
	fmt.Println("Demo: division by 2 in a binary representation.")
	fmt.Printf("%[1]d == %[1]b\n", 8)
	fmt.Printf("%[1]d == %[1]b\n", 8>>1)
	fmt.Printf("%[1]d == %[1]b\n", 8>>2)
	fmt.Printf("%[1]d == %[1]b\n", 8>>3)

	fmt.Println()
	fmt.Printf("%[1]d == %[1]b\n", 7)
	fmt.Printf("%[1]d == %[1]b\n", 7>>1)

	var d uint64 = 39
	fmt.Println("PopCount")
	i := popcount.PopCount(d)
	fmt.Printf("%d has %d ones in a binary representation: %b\n", d, i, d)

	fmt.Println("PopCount2")
	i = popcount.PopCount2(d)
	fmt.Printf("%d has %d ones in a binary representation: %b\n", d, i, d)

	fmt.Println("PopCount3")
	i = popcount.PopCount3(d)
	fmt.Printf("%d has %d ones in a binary representation: %b\n", d, i, d)

	d = 18446744073709551615 // uint64 max value
	i = popcount.PopCount3(d)
	fmt.Printf("%d has %d ones in a binary representation: %b\n", d, i, d)

	fmt.Println("PopCount4")
	d = 39
	i = popcount.PopCount4(d)
	fmt.Printf("%d has %d ones in a binary representation: %b\n", d, i, d)

	d = 18446744073709551615 // uint64 max value
	i = popcount.PopCount4(d)
	fmt.Printf("%d has %d ones in a binary representation: %b\n", d, i, d)
}
