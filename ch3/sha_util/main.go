package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var shaFlag = flag.Int("n", 256, "help message for flag n")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: main [-n 256|384|512 STRING")
		os.Exit(1)
	}
	target := strings.Join(args, " ")
	switch *shaFlag {
	case 256:
		fmt.Printf("%#x\n", sha256.Sum256([]byte(target)))
	case 384:
		fmt.Printf("%#x\n", sha512.Sum384([]byte(target)))
	case 512:
		fmt.Printf("%#x\n", sha512.Sum512([]byte(target)))
	default:
		fmt.Println("Choose correct sha: 256, 384, 512")
		os.Exit(1)
	}
}
