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

// a func can return diff types of results
func isEven(numberr int) (bool, error) {
	if numberr%2 == 0 {
		return true, nil
	} else {
		return false, fmt.Errorf("odd number")
	}
}

func main() {
	PrintMessage("eeeepy")
	text, num := MultipleReturnValues()
	fmt.Printf("the returned values are %s and %d\n", text, num) //multiple return values

	check, err := isEven(2)
	fmt.Printf("%t, %v", check, err)

}
