package testing

import (
	"testing"
	"unicode/utf8"
)

func TestEncode(t *testing.T) {
	s := "hello，万志军"
	t.Log(len(s))
	t.Log(utf8.RuneCountInString(s))

	for i, v := range s {
		t.Log(i, v)
	}

	r := []byte(s)
	for i, v := range r {
		t.Log(i, v)
	}
}
