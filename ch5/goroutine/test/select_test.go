package test

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	c1 := make(chan int)
	go func() {
		time.Sleep(1 * time.Millisecond)
		t.Log("go routine c1")
		c1 <- 1
	}()
	c2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Millisecond)
		t.Log("go routine c2")
		c2 <- 2
	}()
	time.Sleep(5 * time.Millisecond)
	select {
	case n := <-c1:
		t.Log(n)
	case n := <-c2:
		t.Log(n)
		//default:
		//	t.Log("no data")
	}
}

// 隐士代码块
func TestCodeScope(t *testing.T) {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		t.Log(a, b, c)
	}

	if d := 10; true {
		t.Log(d)
	}
}

func TestJsonHandle(t *testing.T) {
	str := `{"data":[{"name":"aaa","age":10},{"name":"bbb","age":20},{"name":"bbb","age":30}]}`
	obj := struct {
		Data []struct {
			Name string `json:"name""`
			Age  int    `json:"age""`
		} `json:"data"`
	}{}
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", obj)
	t.Log(obj.Data[2].Name)
}
