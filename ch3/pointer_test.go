package ch3

import "testing"

func TestPointer(t *testing.T) {
	var a int = 2
	var b *int
	b = &a
	t.Log(b)
	t.Log(*b)
	t.Log(a)
	*b = 10
	t.Log(a)
}
