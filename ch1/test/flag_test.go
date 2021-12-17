package test

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

func TestFalgVar(t *testing.T) {
	name := flag.String("name", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}

func TestPtr1(t *testing.T) {
	name := "lemonwan"
	var ptr *string
	ptr = &name
	t.Log(ptr)
	t.Log(*ptr)
}

func TestPtr2(t *testing.T) {
	var ptr *int
	t.Log(ptr)
	if ptr == nil {
		t.Log("ptr is null")
	} else {
		t.Log(*ptr)
	}
}

/** 指针数组 */
func TestArrPtr(t *testing.T) {
	var ptr [3]*int
	var arr [3]int = [3]int{10, 20, 30}

	for i := 0; i < 3; i++ {
		ptr[i] = &arr[i]
	}

	t.Log(ptr)
	for _, item := range ptr {
		t.Log(item)
		t.Log(*item)
	}
}

/** 指向数组的指针 */
func TestPtrPtr(t *testing.T) {
	var count int = 10
	var ptr *int
	var pptr **int
	ptr = &count
	pptr = &ptr
	t.Log(count)
	t.Log(ptr)
	t.Log(*ptr)
	t.Log(pptr)
	t.Log(**pptr)
}

type Books struct {
	id     int
	title  string
	price  float32
	author string
}

/** 数组存储同一类型的数据集合，结构体存储不同数据类型的集合，相当于Java中的类 */
func TestStruct1(t *testing.T) {
	bookA := Books{1000, "《java study》", 99.8, "lemonwan"}
	bookB := Books{id: 1001, title: "《go study》", price: 80.9, author: "lemonwan"}
	t.Log(bookA)
	t.Log(bookB)
}

func TestStruct2(t *testing.T) {
	var bookA Books
	var bookB Books
	bookA.id = 1003
	bookA.title = "<java study>"

	bookB.id = 1004
	bookB.title = "<go study>"
	t.Log(bookA)
	t.Log(bookA.author)
	t.Log(bookB)
}

func TestStructPtr(t *testing.T) {
	var book Books = Books{1000, "《java study》", 99.8, "lemonwan"}
	showBook(&book, t)
}

func showBook(book *Books, t *testing.T) {
	t.Log(*book)
	t.Log(book.id)
	t.Log(book.title)
	t.Log(book.price)
	t.Log(book.author)
}

func TestMap(t *testing.T) {
	var m1 map[string]string
	m1 = make(map[string]string, 10)
	m1["title"] = "Java study"
	m1["author"] = "lemonwan"
	for key, item := range m1 {
		t.Log(key)
		t.Log(item)
	}

	if data, ok := m1["price"]; !ok {
		t.Log("m1 field not exists")
	} else {
		t.Log(data)
	}
}

func TestGoroutine(t *testing.T) {
	go hello("hello")
	hello("world")
}

func hello(s string) {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestChan(t *testing.T) {
	ch := make(chan int)
	go sum([]int{1, 2, 3}, ch)
	go sum([]int{4, 5, 6}, ch)
	x, y := <-ch, <-ch
	t.Log(x)
	t.Log(y)
	t.Log(x + y)
}

func sum(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}
