package testing

import "testing"

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
