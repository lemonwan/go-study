package main

import (
	"fmt"
	"time"
)

// 定义结构体
type Books struct {
	name  string
	price float32
}

// 定义接口
type paper interface {
	getName() string
	getPrice() float32
}

// 实现方法
func (b Books) getName() string {
	return b.name
}

func (b Books) getPrice() float32 {
	return b.price
}

/** goroutine 执行方法 */
func Print(str string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

// 通道
func fob(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// func main()  {
//   go Print("hello");
//   Print("world")

//   // 调用方法
//   book := Books{"Java入门到实战", 89.9}
//   name := book.getName()
//   price := book.getPrice()
//   fmt.Println(name)
//   fmt.Println(price)

//   // channel
//   var ch chan int = make(chan int, 10)
//   fob(cap(ch), ch)
//   for i := range ch {
//     fmt.Println(i)
//   }
// }
