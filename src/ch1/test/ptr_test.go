package ptr_test

import (
	"testing"
)

func TestPtr1(t *testing.T) {
	t.Log("hello")
	var data int = 10
	// var ptr *int
	// ptr = &data
	ptr := &data
	t.Log(ptr)
	t.Log(*ptr)
}

const LIMIT = 30

func TestPtr2(t *testing.T) {
	fn := fib()
	var arr [LIMIT]int
	for i := 0; i < LIMIT; i++ {
		arr[i] = fn()
	}
	t.Log(arr)
}

func fib() func() int {
	back1, back2 := 0, 1
	return func() int {
		back1, back2 = back2, (back1 + back2)
		return back1
	}
}
