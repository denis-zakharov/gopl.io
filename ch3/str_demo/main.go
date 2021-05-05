package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Привет, строка!" // 15 runes/code points

	fmt.Println("===len vs utf8.RuneCountInString===")
	fmt.Printf("len=%d, rune=%d\n", len(s), utf8.RuneCountInString(s))

	fmt.Println("===UTF-8 Decoder===")
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%#x\t%c\n", i, r, r)
		i += size
	}

	fmt.Println("===Builtint range UTF-8 Decoder===")
	for i, r := range "Hello, Мир!" {
		fmt.Printf("%d\t%c\t%#x\n", i, r, r)
	}

	fmt.Println("==String こんにちは世界 to Rune and back==")
	s = "こんにちは世界"
	fmt.Printf("UTF-8 HEX: % #x\n", s)
	r := []rune(s)
	fmt.Printf("UTF-32 HEX: %#x\n", r)
	fmt.Printf("UTF-32 --> UTF-8: %s\n\n", string(r))

	fmt.Println("==string converts int as a rune code (Unicode)==")
	var c int32 = 0x3053
	fmt.Println(c, string(c)) // an integer is interpreted as a rune in a string constructor
	fmt.Println(12345, string(int32(12345)))
	var undefUTF8 int32 = 1234567
	//var undefReplacement int32 = '\uFFFD'
	fmt.Printf("\nA rune %d is undefined in Unicode. Such runes are replaced by a special symbol %[2]s : %#x\n", undefUTF8, string(undefUTF8), []rune(string(undefUTF8)))
}
