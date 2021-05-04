package main

import (
	"fmt"
)

func main() {
	var ui uint64 = 5
	fmt.Printf("Unsigned: %d == %b\n", ui, ui)

	var si int64 = -5
	fmt.Printf("Signed: %d == %b\n", si, si)

	fmt.Printf("Signed -5 in two's complement representation: %b\n", uint64(si))

	// Manual two's complement conversion 5 --> -5:
	// invert
	ui ^= 0xffffffffffffffff
	// plus one
	ui += 1
	fmt.Printf("Unsigned 5 converted to signed -5: %d == %b\n", int64(ui), ui)
}
