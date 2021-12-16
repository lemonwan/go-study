package main

import (
	"testing"
	"time"
	"fmt"
)

func oneTask() {
	fmt.Println("one task running")
	time.Sleep(time.Second * 2)
}

func twoTask() int{
	time.Sleep(time.Second * 5)
	fmt.Println("two task running")
	return 5
}


func wrapperFunc() chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("LOG:task start run")
		ret := twoTask()
		ch <- ret
		fmt.Println("LOG:task end")
	} ()
	return ch
}

func TestChannel(t *testing.T) {
	ch1 := wrapperFunc()
	oneTask()
	t.Log(<- ch1)
}
