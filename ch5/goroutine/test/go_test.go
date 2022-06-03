package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	now := time.Now()
	t.Log(now)
	t.Log(runtime.NumCPU())
	t.Log(time.Since(now))

	st, err := time.ParseDuration("-24h")
	add := time.Now().Add(st)
	t.Log(add, err)
	date, month, day := add.Date()
	t.Log(date, month.String(), day)
}

func TestChan(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go send(ch)
	go receive(ch, &wg)
	wg.Wait()
}

func send(c chan<- int) {
	c <- 10
}

func receive(c <-chan int, w *sync.WaitGroup) {
	defer w.Done()
	t := <-c
	fmt.Println(t)
}
