package main

import (
	"fmt"
	"unicode/utf8"
)

func NoRepeating(str string) int {
	lastAppear := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, item := range []byte(str) {
		if pos, ok := lastAppear[item]; ok && pos >= start {
			start = pos + 1
		}
		if i-start >= maxLength {
			maxLength = i - start + 1
		}
		lastAppear[item] = i
	}
	return maxLength
}

func HandleString(str string) {
	for _, s := range str {
		fmt.Printf("%X\n", s)
	}
	fmt.Println()
	byteStr := []byte(str)
	for _, s := range byteStr {
		fmt.Printf("%X ", s)
	}
	fmt.Println()
	for i, s := range []rune(str) {
		fmt.Printf("(%d %c)", i, s)
	}
	fmt.Println()
	for len(byteStr) > 0 {
		decodeRune, size := utf8.DecodeRune(byteStr)
		fmt.Printf("%c ", decodeRune)
		byteStr = byteStr[size:]
	}
	fmt.Println(utf8.RuneCount([]byte(str)))
}

func main() {
	fmt.Println(NoRepeating("hello"))
	HandleString("你好world!")
}
