package main

import (
	"sync"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	sign := make(chan struct{}, 3)
	max := 5

	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			t.Logf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			t.Logf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			t.Logf("receiver [%d]: the mailbox is full.", i)
			mailbox = 0
			t.Logf("receiver [%d]: the letter has been received.", i)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
