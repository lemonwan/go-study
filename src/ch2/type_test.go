package main

import "testing"

func TestType(t *testing.T) {
	arr := []string{"zero", "one", "two"}
	/** 类型断言表达式，interface{}代表空接口，任何类型都是它的实现类型 */
	if v, ok := interface{}(arr).([]string); ok == true {
		t.Log(v)
	}
}

func TestString(t *testing.T) {
	str := "你"
	t.Log(len(str))
	t.Logf("%x\n", str)
	// t.Logf("%c\n", str)

	arr := []rune(str)
	t.Log(arr)
	for _, v := range arr {
		t.Logf("%x\n", v)
		t.Log(v)
	}

	for _, v := range str {
		t.Log(v)
	}
}
