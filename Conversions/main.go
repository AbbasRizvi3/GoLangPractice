package main

import (
	"fmt"
	"strconv"
)

func main() {

	// int to string
	x := 5
	s := strconv.Itoa(x)
	fmt.Printf("%s\n", s)

	//string to int
	s2 := "hello"
	x2, _ := strconv.Atoi(s2)
	fmt.Printf("%d\n", x2)

	var ss string = "asd"
	ss = "sdasd"
	fmt.Printf("%s\n", ss)
	//ss[1]='1'// immutable but reassignable as it makes the variable point to another address

}
