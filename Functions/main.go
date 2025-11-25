package main

import "fmt"

// definition of a simple function, there is no void, so write nothing in case of void
// First capital letter of func means function is exportable to other packages
func PrintMessage(message string) {
	fmt.Printf("the message is %s\n", message)
}

func MultipleReturnValues() (string, int) {
	return "sampletext", 0
}

func main() {
	PrintMessage("eeeepy")
	text, num := MultipleReturnValues()
	fmt.Printf("the returned values are %s and %d\n", text, num) //multiple return values

}
