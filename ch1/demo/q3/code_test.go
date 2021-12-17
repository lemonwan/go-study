package test

import (
	// "fmt"
	"testing"
)

func TestCode(t *testing.T) {
	str := "万"
	t.Logf("=> runes(char)：%q\n", []rune(str))
	t.Logf("=> runes(hex)：%x\n", []rune(str))
	t.Logf("=> bytes(hex)：[% x]\n", []byte(str))
}

func TestMap(t *testing.T) {
	m := make(map[int]int, 10)
	t.Log(m[1])
	t.Logf("m length %d", len(m))
	m[1] = 0
	if v, ok := m[1]; ok {
		t.Logf("m[1] = %d", v)
	} else {
		t.Logf("m[1] is not exist")
	}
}

func TestAdvanceMap(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
}
