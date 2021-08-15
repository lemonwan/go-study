package main

import (
	"ch1/demo/q2/lib"
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "greeting to someone")
}

func main() {
	// a := 10
	flag.Parse()
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	fmt.Printf("hello, %s!\n", name)
	lib1.Println(name)
}
