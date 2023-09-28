package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (*Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "%q = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 %s", r.URL.Path)
	}

}

func main() {
	engin := new(Engine)
	log.Fatal(http.ListenAndServe(":8085", engin))

}
