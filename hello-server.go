package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, sunshine!\n")
}

func main() {
	fmt.Printf("Env var: %s = %s \n", "environment", os.Getenv("environment"))
	http.HandleFunc("/", HelloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
