package main

import (
	"crypto/sha256"
	"fmt"
	"gopl.io/denis-zakharov/ch2/popcount"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: script string1 string2")
		os.Exit(1)
	}
	c1 := sha256.Sum256([]byte(args[0]))
	c2 := sha256.Sum256([]byte(args[1]))
	fmt.Printf("%s\n%#x\n%#b\n\n%s\n%#x\n%#b\n\n", args[0], c1, c1, args[1], c2, c2)

	fmt.Println("Number of different bits sha_1 XOR sha_2")
	res := 0
	for i := range c1 {
		diff := c1[i] ^ c2[i]
		res += popcount.PopCount4(uint64(diff))
		fmt.Printf("%b ", diff)
	}
	fmt.Printf("\nDiff bits: %d\n", res)
}
