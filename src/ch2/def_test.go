package main

import (
	"errors"
	"testing"
)

func TestDef(t *testing.T) {
	t.Log("Enter testing func")
	defer func() {
		t.Log("Enter defer func")
		if p := recover(); p != nil {
			t.Logf("panic: %s\n", p)
		}
		t.Log("Exit defer func")
	}()
	panic(errors.New("something wrong"))
	t.Log("Exit testing func")
}
