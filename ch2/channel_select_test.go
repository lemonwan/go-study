package main

import (
	"fmt"
	"testing"
	"time"
)

func oneWork() chan string {
	ch := make(chan string)
	go func() {
		fmt.Println("one work staring run")
		time.Sleep(time.Second * 50)
		ch <- "end"
		fmt.Println("one work end")
	}()
	return ch
}

func twoWork() {
	fmt.Println("two work staring run")
	time.Sleep(time.Second * 10)
	fmt.Println("two work end")
}

func TestSelectChannel(t *testing.T) {
	ch := oneWork()
	twoWork()
	select {
	case ret := <-ch:
		t.Log(ret)
		// t.Log("all task run end")
	case <-time.After(time.Millisecond * 30):
		t.Error("timeout")
	}
}
