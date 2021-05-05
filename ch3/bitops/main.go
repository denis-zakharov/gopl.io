package main

import "fmt"

/* Operands of bit operations are interpreted as raw bits without sign.
It is convenient to interpret bits as a set of zero-indexed numbers from
the lowest bit.
*/

func main() {
	var x uint8 = 1 | 1<<1 | 1<<5 // {0, 1, 5}
	var y uint8 = 1<<1 | 1<<2     // {1, 2}

	fmt.Printf("x == %08b\t{0, 1, 5}\n", x)
	fmt.Printf("y == %08b\t{1, 2}\n", y)
	fmt.Printf("Intersection x&y == %08b\t{1}\n", x&y)
	fmt.Printf("Union x|y == %08b\t{0, 1, 2, 5}\n", x|y)
	fmt.Printf("Symmetric diff x^y == %08b\t{0, 2, 5}\n", x^y)
	fmt.Printf("Difference %08b\t{0, 5}\n", x&^y)

	fmt.Printf("\nRight shift of _signed_ numbers populates higher bits with copies of a sign bit.\n")
	var signed int8 = -127
	fmt.Printf("%08b\n", uint8(signed))
	fmt.Printf("%08b\n", uint8(signed>>2))
}
