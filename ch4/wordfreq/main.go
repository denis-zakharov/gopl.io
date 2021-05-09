package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main filename(s)")
		os.Exit(1)
	}

	counts := make(map[string]map[string]int) // filename -> word -> count
	for _, filename := range os.Args[1:] {

		fileCounts := make(map[string]int)
		counts[filename] = fileCounts

		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "count error: %v\n", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			fileCounts[scanner.Text()]++
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "count error: %v\n", err)
		}
	}

	for filename := range counts {
		fmt.Printf("===%s===\n", filename)
		fmt.Println(counts[filename])
		fmt.Println()
	}

}
