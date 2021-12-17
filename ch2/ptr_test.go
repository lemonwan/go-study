package main

import (
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Dog struct {
	name string
	age  int
}

func TestStrucPtr(t *testing.T) {
	dog := Dog{"pig", 10}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	t.Log(namePtr)
	ptrP := (*string)(unsafe.Pointer(namePtr))
	t.Log(ptrP)
	t.Log("--------")
	name := dogP.name
	t.Log(&name)
	t.Log(&(dogP.name))
}

func TestPtr(t *testing.T) {
	i := 10
	ptr := &i
	t.Log(ptr)

	arr := [3]int{12, 13, 14}
	t.Log(&(arr[0]))
}

/** 测试goroutine运行情况：如下程序输出结果？ */
func TestGoRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			t.Log(i)
		}()
	}
}

func TestGoRoutineNew(t *testing.T) {
	num := 10
	sign := make(chan struct{}, num)
	for i := 0; i < 10; i++ {
		go func() {
			t.Log(i)
			sign <- struct{}{}
		}()
	}

	for j := 0; j < num; j++ {
		<-sign
	}
}

/** 控制goroutine顺序执行输出 */
func TestGoroutineOrderExec(t *testing.T) {
	var count uint32

	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				t.Log(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}

func TestFor(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	for i := range s1 {
		if i == 3 {
			s1[i] |= i
		}
	}
	t.Log(s1)
	t.Log("------------------")

	numbers := []int{1, 2, 3, 4, 5, 6}
	maxIndex := len(numbers) - 1
	for i, e := range numbers {
		if i == maxIndex {
			numbers[0] += e
		} else {
			numbers[i+1] += e
		}
	}
	t.Log(numbers)
}
