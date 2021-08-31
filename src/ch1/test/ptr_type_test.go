package ptr_test

import "testing"

func TestType(t *testing.T) {
	var ptr *int
	i := 10
	ptr = &i
	t.Log(ptr)
	t.Log(*ptr)
}
