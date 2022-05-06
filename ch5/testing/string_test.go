package testing

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestEncode(t *testing.T) {
	s := "hello，万志军"
	t.Log(len(s))
	t.Log(utf8.RuneCountInString(s))

	fmt.Println("-------")
	for i, v := range s {
		t.Log(i, v)
		t.Log(string(v))
	}

	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		ch := fmt.Sprintf("%c", r)
		t.Log(ch)
	}

}
