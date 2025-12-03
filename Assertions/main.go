package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct{ Name string }

func (d Dog) MakeSound() string { return "Woof!" }

type Cat struct{ Name string }

func (c Cat) MakeSound() string { return "Meow!" }

func main() {

	animal := []Animal{Dog{Name: "bad doggie"}, Cat{Name: "2"}}
	fmt.Printf("%s\n", animal[0].(Dog).Name)
	//assertion is needed as compiler can only access interfaces methods, not methods and values in actual type

	var i interface{} = "hello" //empty interface and can store any time
	slicess := make([]interface{}, 0, 10)
	slicess = append(slicess, 1)
	slicess = append(slicess, "2")
	fmt.Printf("%v\n", slicess...)
	// the use of interface helps store any value as empty interface allows that
	fmt.Printf("%T\n", i)
	i = 2
	fmt.Printf("%T\n", i)
	x := i.(int)
	x = x + 1
	fmt.Printf("%d\n", x)

	// this way is risky as assertion means that we force that a certain value exist
	// i = 1
	// temp := i.(string)
	// fmt.Printf("%v\n", temp)
	// this will result in an error because we specified that it is of type string when it isnt

	// safe assertion
	i = 1
	temp, ok := i.(string)
	if ok {
		fmt.Printf("%v\n", temp)
	} else {
		fmt.Printf("failed\n")
	}

	//this way the program wont crash

	//there is smth called a switch type too

	switch i.(type) {
	case string:
		fmt.Print("string")
	case int:
		fmt.Print("int")
	}
}
