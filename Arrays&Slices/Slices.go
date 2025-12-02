package main

import "fmt"

func ModifySlice(array []int) {
	array = append(array, 9)
}

func main() {
	// slices are part of arrays. they share the same underlying memory so a change in either of the values reflect changes on both sides

	// there are slices of an array and not slices itself

	//slices are pass by ref by default but may result in pass by value in some cases
	array := [5]int{1, 2, 3, 4, 5}
	sliced_array := array[2:]

	fmt.Printf("%d\n", array)
	fmt.Printf("%d\n", sliced_array)

	sliced_array[2] = 99
	array[0] = 45

	fmt.Printf("%d\n", array)
	fmt.Printf("%d\n", sliced_array)

	//Slices

	slice_array := []int{1, 2, 3, 4, 5}
	//slice_array2:=make([]int,3,5) makes an array of size 3 and max size(flexibility) of 5

	slice_array = append(slice_array, 66)
	fmt.Printf("%d\n", slice_array)

	arrayyy := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sliced_arrayyy := arrayyy[:5]
	fmt.Printf("%d\n", sliced_arrayyy)

	// add an element at a specific index
	arrayyy = append(arrayyy[:5], append([]int{99}, arrayyy[5:]...)...)
	fmt.Printf("%d\n", arrayyy)

	//removing an element at a specific index
	arrayyy = append(arrayyy[:5], arrayyy[6:]...)
	fmt.Printf("%d\n", arrayyy)

	temp_slice := make([]int, 5)
	ModifySlice(temp_slice)
	fmt.Printf("%d\n", temp_slice)

}
