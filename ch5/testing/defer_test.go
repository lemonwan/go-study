package testing

import (
	"bufio"
	"ch5/functional/fib"
	"fmt"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	writeFile("defer.txt")
}

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		var err = file.Close()
		if err != nil {

		}
	}(file)
	writer := bufio.NewWriter(file)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {

		}
	}(writer)
	f := fib.F()
	for i := 0; i <= 20; i++ {
		_, err := fmt.Fprintln(writer, f())
		if err != nil {
			return
		}
	}
}
