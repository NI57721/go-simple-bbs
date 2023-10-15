package main

import (
	"io"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	go main()

	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var got, want string

	want = "Hello, world!\n"
	got = string(body)
	if got != want {
		t.Errorf("The body of http.Get(\"localhost:8080/hello\") is %v, but wants %v", got, want)
	}
}
