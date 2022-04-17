package main

import (
	"fmt"
	"sync"
)

func createWorker(i int, wg *sync.WaitGroup) worker {
	work := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(i, work)
	return work
}

func doWork(i int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", i, n)
		w.wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func main() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		workers[i].in <- 'A' + i
	}

	wg.Wait()
}
