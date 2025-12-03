package main

import "fmt"

type Person[T int | float32, F int | float32] struct {
	Name T
	Age  F
}

type Values interface {
	int | float64
}

// generics support multiple parameter types for functions

// this will only work for 2 ints
func add(a int, b int) int {
	return a + b
}

// a generic function

func Add[V int | float32 | float64](a V, b V) V {
	return a + b
}

// there is a comparable type that includes only comparable operations
// int float strings arrays structs and doesnt include slices maps functions

func comparison[T comparable](X T, Y T) bool {
	return X == Y
}

func IntorFloatMapSum[T comparable, F Values](m map[T]F) F {
	var sum F
	for _, value := range m {
		sum += value
	}
	return sum
}

// type alias
type MYINT = int

//it means MYINT and int are the same thing

// but what if i do this
type MYINT2 int

// now underline they are the same thing but actually are not completely diff types

type MYINT3 interface {
	~int
}

func Add2[V MYINT3](a V, b V) V {
	return a + b
}

func main() {
	var x, y MYINT
	x = 3
	y = 4
	fmt.Printf("%d\n", Add(x, y))
	// it works as they are the same thing because of =

	var q, e int
	q = 3
	e = 4
	// fmt.Printf("%d\n",Add(q,e))
	//this doesnt work as myint2 is diff from int

	//to make it work, use ~ notation
	fmt.Printf("%d\n", Add(q, e))

	fmt.Printf("%d\n", add(5, 5))
	fmt.Printf("%d\n", Add(5, 5))
	fmt.Printf("%v\n", Add(5.5, 5.5))
	fmt.Printf("%v\n", Add(5, 7.5))

	fmt.Printf("%t\n", comparison(5, 3))
	fmt.Printf("%t\n", comparison("a", "a"))

	// allows generic operations

	map1 := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	map2 := map[string]float64{
		"1": 3.1,
		"2": 5.6,
		"3": 7.9,
	}

	fmt.Printf("%d\n", IntorFloatMapSum(map1))
	fmt.Printf("%f\n", IntorFloatMapSum(map2))

}
