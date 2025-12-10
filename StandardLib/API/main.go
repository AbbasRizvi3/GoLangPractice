package main

import (
	"fmt"
	"net/http"
)

func handlerf(h http.ResponseWriter, r *http.Request) {
	fmt.Printf("The request type is: %v\n", r.URL.RequestURI())
	fmt.Printf("The header is: %v\n", r.Header)

	h.WriteHeader(200)
	fmt.Fprint(h, "Welcome")
}

func main() {
	http.HandleFunc("/", handlerf)
	http.ListenAndServe(":8000", nil)
}
