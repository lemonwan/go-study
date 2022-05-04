package main

import (
	"ch5/api"
	"fmt"
)

// 接口组合
const url = "https://www.google.com"

type R interface {
	Get(url string) string
}

type B interface {
	Post(url string, param map[string]string) string
}

type C interface {
	R
	B
}

func getObj() C {
	return &api.Receiver{}
}

func post(c C) string {
	return c.Post(url, map[string]string{"contents": "hello"})
}

func main() {
	receiver := &api.Receiver{Contents: "main"}
	var c C = receiver
	switch r := c.(type) {
	// 第一种：判断接口类型的方式
	case *api.Receiver:
		fmt.Println(r.Contents)
	default:
		fmt.Println("non")
	}

	fmt.Printf("%T %v \n", c, c)
	// 第二种：判断接口类型的方式
	if e, ok := c.(*api.Receiver); ok {
		fmt.Println("type assertion", e.Contents)
	}

	s := post(c)
	fmt.Println(s)
	post(receiver)
	fmt.Println(receiver.Contents)
	fmt.Println("*************")
	fmt.Println(receiver)
}
