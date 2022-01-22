package ch4

import (
	"io/ioutil"
	"math"
	"strconv"
	"testing"
)

const file = "abc.txt"

func TestConst(t *testing.T) {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	t.Log(c, file)
}

func TestConst2(t *testing.T) {
	const (
		java = iota
		_
		php
		javascript
		python
	)
	t.Log(java, php, javascript, python)
}

func TestBranch(t *testing.T) {
	const filename = "../README.md"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(contents)
		t.Logf("%s\n", contents)
	}
	// 简单写法：但是在if代码块以外不能使用if内定义的变量
	if con, err := ioutil.ReadFile(filename); err != nil {
		t.Log(err)
	} else {
		t.Logf("%s\n", con)
	}
}

func TestConvertToBin(t *testing.T) {
	t.Log(convertToBin(5))
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result += strconv.Itoa(lsb)
	}
	return result
}

func TestMap(t *testing.T) {
	var m map[string]string       // m == nil
	m1 := make(map[string]string) // m1 == empty map
	t.Log(m, m1)
	if m == nil {
		t.Log("m is nil")
		t.Log(len(m))
	}
	if m1 == nil {
		t.Log("m1 is nil")
	}
}
