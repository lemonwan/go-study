package ch3

import (
	"reflect"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("create new object")
			return 100
		},
	}
	v := pool.Get().(int)
	t.Log(v)

	pool.Put(3)
	runtime.GC()
	v1 := pool.Get().(int)
	t.Log(v1)

	v2 := pool.Get().(int)
	t.Log(v2)
}

func TestReflect(t *testing.T) {
	a := 10
	ref := reflect.TypeOf(a)
	t.Log(ref)
	b := reflect.ValueOf(a)
	t.Log(b)
	t.Log(reflect.ValueOf(a).Type())
	t.Log(reflect.Float32)
	t.Log(ref.Kind())
}
