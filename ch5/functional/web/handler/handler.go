package handler

import (
	"io/ioutil"
	"net/http"
	"os"
)

func List(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(contents)
	return nil
}
