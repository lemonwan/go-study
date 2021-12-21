package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func taskDone(i int) string {
	time.Sleep(time.Millisecond * 5)
	return fmt.Sprintf("task running from %d", i)
}

func goTask() string {
	task := 10
	ch := make(chan string)
	for i := 0; i < task; i++ {
		go func(i int) {
			ret := taskDone(i)
			ch <- ret
		}(i)
	}
	return <-ch
}
func TestTask(t *testing.T) {
	t.Log(goTask())
	time.Sleep(time.Millisecond * 1)
}
