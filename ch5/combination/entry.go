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
	var c C
	receiver := &api.Receiver{Contents: "main"}
	c = receiver
	s := post(c)
	fmt.Println(s)
	post(receiver)
	fmt.Println(receiver.Contents)
	fmt.Println("*************")
	fmt.Println(receiver)
}
