package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 生产者
func producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()
}

// 消费者
// 一对一的生产和消费，可以保证程序正常运行
// 如果是一生产对多个消费，则会出现程序卡主的问题，通道阻塞
func consumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	producer(ch, &wg)
	wg.Add(1)
	consumer(ch, &wg)
	wg.Wait()
}

// 改进程序
func newProducter(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func newConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelCloseNew(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	newProducter(ch, &wg)
	wg.Add(1)
	newConsumer(ch, &wg)
	wg.Add(1)
	newConsumer(ch, &wg)
	wg.Wait()
}

func isCancel(ch chan struct{}) bool {
	select {
	// 通道如果关闭了，继续向通道发送数据，则会产生panic错误
	// 通道如果关闭了，继续从通道获取数据，会立即得到通道数据类型的默认零值
	case <-ch:
		return true
	default:
		return false
	}
}

func cancelOne(ch chan struct{}) {
	close(ch)
}

func cancelTwo(ch chan struct{}) {
	ch <- struct{}{}
}
func TestTaskCancel(t *testing.T) {
	ch := make(chan struct{}, 0)
	for i := 0; i < 10; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancel(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
				t.Log(i, "run……")
			}
			t.Log(i, "success cancel")
		}(i, ch)
	}
	cancelOne(ch)
	time.Sleep(time.Millisecond * 1)
}
