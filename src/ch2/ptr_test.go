package main

import (
	"testing"
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
