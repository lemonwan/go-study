package testing

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func TestOpen(t *testing.T) {
	open, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	}
	printString(open, t)
}

func TestPrintString(t *testing.T) {
	str := `1234

abc
qwe`
	reader := strings.NewReader(str)
	printString(reader, t)
}

func printString(r io.Reader, t *testing.T) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
}
