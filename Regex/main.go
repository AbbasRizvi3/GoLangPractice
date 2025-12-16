package main

import (
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed File.txt
var message string

func main() {
	text := "Hello, welcome to new york"
	regex := regexp.MustCompile(`H[a-z]{4} | y[A-z]{3}`)
	if regex.MatchString(text) {
		fmt.Print("Contains")
	}

	fmt.Println(message)

}
