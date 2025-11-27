package main

import "fmt"

func main() {
	//delclaring the array only
	// var arr2 [10]int

	//declaring and initializing
	var arr = [10]int{1, 2, 3, 4}
	fmt.Printf("%v\n", arr)

	// arr2:=[10]int{1,3,4,4}
	// var arr2 =[...]int{1,2,3,4} size equals to number of elements

	//array of functions

	var functional_array [2]func(int) int
	functional_array[0] = func(x int) int {
		return x + 1
	}
	functional_array[1] = func(x int) int {
		return x + 2
	}

	fmt.Printf("%d\n", functional_array[0](5))
	fmt.Print("%d\n", functional_array[1](5))
	fmt.Printf("\n")

}
