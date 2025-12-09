package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v\n", r.Method)
	fmt.Fprint(w, "Hey")
	fmt.Printf("Header %v\n", r.Header)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
