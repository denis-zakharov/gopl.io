package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elemcount: %v\n", err)
		os.Exit(1)
	}
	for k, v := range visit(make(map[string]int), doc) {
		fmt.Println(k, v)
	}
}

func visit(elemCounts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elemCounts[n.Data]++
	}
	if sibling := n.NextSibling; sibling != nil {
		elemCounts = visit(elemCounts, sibling)
	}
	if child := n.FirstChild; child != nil {
		elemCounts = visit(elemCounts, child)
	}

	return elemCounts
}
