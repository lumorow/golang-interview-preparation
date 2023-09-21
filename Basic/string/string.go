package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "ご気分はいかがですか?"

	fmt.Println(str)
	for i, c := range str {
		fmt.Printf("%v %c\n", i, c)
	}
	fmt.Println(len(str), "bytes")                    // Output: 31 bytes
	fmt.Println(utf8.RuneCountInString(str), "runes") // Output: 11 runes

	c, size := utf8.DecodeRuneInString(str)
	fmt.Printf("First rune: %c %v bytes", c, size) // Output: First rune: ご 3 bytes
}
