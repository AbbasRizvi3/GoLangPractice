package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	var person1 *person = new(person)
	person1.Age = 20
	person1.Name = "jja"
	fmt.Printf("%s\n", person1.Name)

	var x *int = new(int)
	*x = 43
	fmt.Printf("%d\n", *x)

	s := []int{1, 2, 3, 4, 6, 7}
	s = append(s[:4], s[3:]...)
	s[4] = 5

	fmt.Printf("%d\n", s)

	s = append(s[:5], s[6:]...)
	fmt.Printf("%d\n", s)

}
