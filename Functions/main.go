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

// anonymous function
var anonymous_func = func(x int) int {
	return x * 2
}

// named return values
func NamedReturnFunc() (model string, typeof string) {
	model = "honda"
	typeof = "Sedan"
	return
}

// the use of ... is kind of like spread operator in javascript that can accept multple arguments and acts like a slice in func
func Sum(numbers ...int) int {
	fmt.Printf("%T\n", numbers)
	var total int
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

// here the elements passed to function were treated as a slice
// the ... notation can be used to spread arguments too when calling a func
func Sum2(numbers ...int) int {
	var total int
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

//closure func acts a bit like classes. the outer scope acts as a constructor and its value persists
//in heap instead of stack if it is to be used in closure functions.

func multiplier(multiplyer int) func(int) int {
	n := multiplyer
	func1 := func(x int) int {
		return n * x
	}
	return func1
}

// a function that returns multiple closures
func multiples(x int) (func(int) int, func(int) int) {
	func1 := func(z int) int {
		return x * z
	}
	func2 := func(y int) int {
		return y * x
	}

	return func1, func2
}

func main() {
	PrintMessage("eeeepy")
	text, num := MultipleReturnValues()
	fmt.Printf("the returned values are %s and %d\n", text, num) //multiple return values

	check, err := isEven(2)
	fmt.Printf("%t, %v", check, err)
	fmt.Printf("\n")
	ans := anonymous_func(2)
	fmt.Printf("%d\n", ans)

	model, typeof := NamedReturnFunc()
	fmt.Printf("the model is %s, and the make type is %s\n", model, typeof)

	total := Sum(1, 2, 3, 4, 5)
	fmt.Printf("The total is %d\n", total)

	a := multiplier(2)
	b := multiplier(3)
	fmt.Printf("%d\n", a(2))
	fmt.Printf("%d\n", b(2))

	x, y := multiples(10)
	fmt.Printf("func1: %d, func2: %d \n", x(2), y(3))

	arrayy := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%d\n", Sum2(arrayy[:]...))

}
