package testing

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	data := []struct{ a, b, c int }{
		{1, 2, 3},
		{4, 5, 9},
		{5, 5, 10},
	}

	for _, v := range data {
		if sum := add(v.a, v.b); sum != v.c {
			t.Errorf("add(%d, %d); got %d; expected %d", v.a, v.b, v.c, sum)
		}
	}
}

func TestReader(t *testing.T) {
	reader := strings.NewReader("ABCDEF")
	newReader := bufio.NewReaderSize(reader, 0)
	// b := make([]byte, 10)
	peek, _ := newReader.Peek(10)
	t.Logf("%d == %q\n", newReader.Buffered(), peek)
	fmt.Println("10")
	fmt.Printf("%#v\n", 10)
	fmt.Printf("%v\n", 10)
	fmt.Printf("%.4g\n", 123.45)
	sprintf := fmt.Sprintf("%.4g\n", 123.45)
	t.Log(sprintf)
	fmt.Printf("%6.2f\n", 123.45)
}
