package main

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct{}

var instance *Singleton
var once sync.Once

func getSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		instance = new(Singleton)
	})
	return instance
}

func TestRun(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := getSingletonObj()
			t.Logf("%x", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
