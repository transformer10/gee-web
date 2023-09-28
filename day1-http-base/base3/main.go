package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engin := gee.New()
	engin.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	engin.Get("/tset", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q", r.Header)
	})
	engin.Run(":8085")

}
