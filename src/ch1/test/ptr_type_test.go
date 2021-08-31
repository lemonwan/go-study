package ptr_test

import (
	"container/list"
	"testing"
)

func TestType(t *testing.T) {
	var ptr *int
	i := 10
	ptr = &i
	t.Log(ptr)
	t.Log(*ptr)
}

func TestList(t *testing.T) {
	li := list.New()
	l1 := li.PushBack(4)
	l2 := li.PushFront(1)
	li.InsertAfter(2, l2)
	li.InsertBefore(3, l1)
	for e := li.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
}
