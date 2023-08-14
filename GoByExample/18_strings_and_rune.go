package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// GOlang string is a read-only slice of bytes
	// string is a container of text encoded in utf-8
	// in other languages made of characteres, here rune-integer representing a unicode codepoint

	s := "สวัสดี"

	fmt.Println("String length:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}

	fmt.Println("runeLength:", utf8.RuneCountInString(s))
	for index, codepoint := range s {
		fmt.Printf("%#U starts at %d\n", codepoint, index)
	}

	fmt.Println("using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s)
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}
}
