package main

import (
	"flag"
	"fmt"
)

func main() {
	// name := flag.String("name", "Guest", "Enter your name")
	// age := flag.Int("age", 0, "Enter your age")
	// flag.Parse()
	// fmt.Printf("name: %s\n", *name)
	// fmt.Printf("age: %d\n", *age)

	var name2 string
	var age2 int

	flag.StringVar(&name2, "name", "guest", "Enter your name")
	flag.IntVar(&age2, "age", 0, "Enter your age")

	flag.Parse()
	fmt.Printf("%s\n", name2)
	fmt.Printf("%d\n", age2)

}
