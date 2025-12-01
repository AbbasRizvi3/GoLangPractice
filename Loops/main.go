package main

import "fmt"

func main() {

	// can initialize bfr writing the condition. The initialized variable is available in if and all else branches only
	if x := 9; x < 15 {
		print("true")
	} else {
		print("false")
	}

	array := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(array); i++ {
		fmt.Printf("%d\n", array[i])
	}
	fmt.Printf("\n")

	//we also have a range loop, it returns 2 values, index and value
	for _, val := range array {
		fmt.Printf("%d\n", val)
	}
	fmt.Printf("\n")

	// can u break to break a loop but use tags to break from an outer one
outer:
	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if i == j {
				break outer
			}
			fmt.Printf("i=%d & j=%d\n", i, j)
		}
	}

}
