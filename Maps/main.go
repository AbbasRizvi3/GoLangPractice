package main

import "fmt"

func main() {
	var map1 = make(map[int]string)
	map1[1] = "aa"
	map1[2] = "bb"

	var value, ok = map1[1]

	if ok {
		fmt.Printf("%s\n", value)
	}
	fmt.Printf("%v", map1)
	delete(map1, 2)
	fmt.Printf("%v", map1)

	for _, val := range map1 {
		fmt.Printf("%s\n", val)
	}

}
