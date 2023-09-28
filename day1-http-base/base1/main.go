package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8085", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}
