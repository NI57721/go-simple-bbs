package main

import (
	"fmt"
)

func foo(s string) string {
	return "foo: " + s
}

func main() {
	fmt.Println(foo("bar"))
}
