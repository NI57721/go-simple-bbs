package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Task struct {
	id      int
	body    string
	is_done bool
}

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
