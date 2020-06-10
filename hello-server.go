package main

import (
	"fmt"
	"net/http"
)

// HTTP Handler - Hello, sunshine
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, sunshine!\n")
}

// HTTP Handler - Echo headers into response body
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
