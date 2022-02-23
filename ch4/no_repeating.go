package main

import "fmt"

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

func main() {
	fmt.Println(NoRepeating("hello"))
}
