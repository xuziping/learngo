package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Yes我是谁"

	for _,b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()

	for i,b := range s {
		fmt.Printf("(%d %X) ", i, b)
	}
	fmt.Println()


	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := [] byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i,ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

	fmt.Println(len(s))
}
