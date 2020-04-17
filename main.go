package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!!")
}

func BaBo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Babo!!!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})

	http.HandleFunc("/babo", BaBo)

	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}
