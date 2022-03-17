package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Receiver struct {
	Contents string
}

func (r *Receiver) Post(url string, param map[string]string) string {
	fmt.Println(url)
	r.Contents = param["contents"]
	return "Receiver Post method"
}

func (*Receiver) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
