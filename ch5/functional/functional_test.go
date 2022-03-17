package functional

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func add() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func TestFuncAdd(t *testing.T) {
	a := add()
	for i := 0; i < 10; i++ {
		t.Log(a(i))
	}
}

func f() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestF(t *testing.T) {
	fu := f()
	t.Log(fu())
	t.Log(fu())
	t.Log(fu())
	t.Log(fu())
	t.Log(fu())
	t.Log(fu())
}

type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printString(r io.Reader, t *testing.T) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
}

func TestFuncRead(t *testing.T) {
	fu := f()
	t.Logf("%T", fu)
	printString(fu, t)
}
