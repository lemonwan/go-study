package main

import (
	"ch5/api"
	"fmt"
)

type ReceiverInter interface {
	Get(str string) string
}

func getReceiver() ReceiverInter {
	// return new(api.Receiver)
	return &api.Receiver{Contents: "hello world"}
}

func main() {
	// logger,_ := zap.NewProduction()
	// logger.Warn("zap warning testing")
	receiver := getReceiver()
	str := receiver.Get("https://www.imooc.com")
	fmt.Println(str)
	fmt.Printf("%T %v", receiver, receiver)
	switch receiver.(type) {
	case *api.Receiver:
		fmt.Println("api.Receiver")
	default:
		fmt.Println("unknown")
	}

	if inter, ok := receiver.(*api.Receiver); ok {
		fmt.Println(inter)
	} else {
		fmt.Println("unknown")
	}
}
