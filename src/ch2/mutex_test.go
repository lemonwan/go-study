package main

import (
	"testing"
	"sync"
	"time"
)

// 不使用并发控制，开启多个协程操作同一共享数据
func TestNoMutex(t *testing.T) {
	counter := 0
	for i := 1; i <= 5000; i++ {
		go func() {
			counter++
		} ()
	}
	// 此处之所以要休眠是为了防止当前goroutine执行结束了，但是其开启的goroutine还未执行结束
	time.Sleep(time.Second * 1)
	t.Log(counter)
}

// 使用互斥锁保护共享资源并发操作
func TestMutex(t *testing.T) {
	counter := 0
	// 使用并发控制工具-互斥锁
	var mut sync.Mutex
	for i := 1; i <= 5000; i++ {
		go func() {
			defer func() {
				// 执行完毕要释放锁
				mut.Unlock();
			} ()
			// 每一个goroutine执行任务之前获取到互斥锁
			mut.Lock()
			counter++
		} ()
	}
	// 主goroutine等待其开启的子goroutine执行完毕，但是存在一个问题是等待多久时间是无法预估的
	time.Sleep(time.Second * 1)
	t.Logf("counter result is %d", counter)
}


// 增加等待等待和通知
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 1; i <= 5000; i++ {
		// 任务执行前增加
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			} ()
			mut.Lock()
			counter++
			// 任务执行完毕减少
			wg.Done()
		} ()
	}
	// 解决主goroutine无法预估等待时间的问题
	wg.Wait()
	t.Logf("counter result is %d", counter)
}
