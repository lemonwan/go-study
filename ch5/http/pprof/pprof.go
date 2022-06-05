package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":8080", nil)
}
