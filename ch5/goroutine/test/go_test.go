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

func TestLock(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	t.Log("Locking (G0)")
	mutex.Lock()
	t.Log("Locked (G0)")
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			t.Logf("Locking (G%d)\n", i)
			mutex.Lock()
			t.Logf("Locked (G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			t.Logf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	t.Log("ready unlock (G0)")
	mutex.Unlock()
	t.Log("unlocked (G0)")
	wg.Wait()
}
