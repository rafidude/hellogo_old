package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!"))
}
