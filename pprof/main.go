package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // import pprof
)

func work() {
	sum := 0
	for i := 0; i < 1_000_000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	go func() {
		for {
			work()
		}
	}()

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
