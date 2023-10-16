package main

import (
	"fmt"
	"net/http"
)

func foo(s string) string {
	return "foo: " + s
}

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello from hi1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
          fmt.Fprint(w, "<!DOCTYPE html><html><link rel=\"stylesheet\" type=\"text/css\" href=\"/assets/style.css\"><title>Simple BBS</title><style><!-- .foo{background:yellow;} --></style>out of p<br><p>inside of p, and <span class=\"foo\" style=\"font-weight: bold;\">inside of span</span></p></html>\n")
	}

	css := func(w http.ResponseWriter, _ *http.Request) {
          fmt.Fprint(w, ".foo{color:red;}\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/about", h2)
	http.HandleFunc("/assets/style.css", css)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
