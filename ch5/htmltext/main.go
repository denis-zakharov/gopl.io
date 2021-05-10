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
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode && n.Parent.Data != "script" &&
		n.Parent.Data != "noscript" && n.Parent.Data != "style" &&
		len(n.Data) > 0 {
		fmt.Println(n.Parent.Data)
		fmt.Println(n.Data)
	}
	if sibling := n.NextSibling; sibling != nil {
		visit(sibling)
	}
	if child := n.FirstChild; child != nil {
		visit(child)
	}
}
