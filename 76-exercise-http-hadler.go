package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(s))
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Who, s.Punct, s.Greeting)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gopher!"})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
