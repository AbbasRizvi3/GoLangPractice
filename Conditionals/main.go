package main

import (
	"fmt"
	"sort"
)

func main() {
	//paranthesis are not allowed in conditions and loops and else must be written on the same line of end block of if statement
	if 2 > 3 {
		print("yes")
	} else {
		print("no")
	}

	var arr = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\n", arr[i])
	}

	//range return 2 values, index and the value

	for _, v := range arr {
		fmt.Printf("value %d\n ", v)
	}

	print("\n")

	//array slicing, [inclusive,exclusive] default=[0,end]
	fmt.Printf("%d\n", arr[3:5])

	// == operator, 2 arrays are same if length is same and elements are in the same order

	ar1 := [3]int{1, 2, 3}
	ar2 := [3]int{1, 2, 3}
	ar3 := [3]int{1, 2, 4}
	fmt.Printf("%t\n", ar1 == ar2)
	fmt.Printf("%t\n", ar1 == ar3)

	//sorting

	sort_an_array := [5]int{4, 6, 7, 5, 1}

	sort.Ints(sort_an_array[:])
	fmt.Printf("%d\n", sort_an_array)

	s_array := [3]string{"d", "b", "c"}

	sort.Strings(s_array[:])
	fmt.Printf("%s\n", s_array)

	//copying arrays
	var copy_array [5]string
	copy(copy_array[:], s_array[:])
	fmt.Printf("%s\n", copy_array)

	//searching a element in array

	students := [4]string{"Feeehu", "Abbas", "Mamoona", "Roha"}
	sort.Strings(students[:])

	index := sort.Search(len(students), func(i int) bool {
		if students[i] == "Roha" {
			return true
		} else {
			return false
		}
	})

	fmt.Printf("found at %d\n", index)

	arrayy := [5]int{45, 2, 32, 3, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(arrayy[:])))

	array_of_slices := [3][]int{
		{1, 2, 3},
		{4},
		{5, 6},
	}
	for i := 0; i < len(array_of_slices); i++ {
		for j := 0; j < len(array_of_slices[i]); j++ {
			fmt.Printf("%d", array_of_slices[i][j])
		}
		print("\n")
	}

	// can initialize bfr writing the condition. The initialized variable is available in if and all else branches only
	if x := 9; x < 15 {
		print("true")
	} else {
		print("false")
	}

	fmt.Printf("\n")
	var user_input int
	fmt.Scanf("%d", &user_input)
	switch user_input {
	case 1:
		fmt.Print("1")
	case 2:
		fmt.Print("2")
	default:
		fmt.Print("0")
	}

	//we can use fallthrough to also execute the next statement only too if a match is found

}
