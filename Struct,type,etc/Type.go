package main

import "fmt"

// we can create alias with types. an important usecase of it is wth use of closure functions to setup default parameters
// along with overloaded ones when making a constructor as overloaded functions, constructors are not available in go

// for example consider a struct called a person

type Person struct {
	Name       string
	Age        int
	Address    string
	Gender     rune
	BloodGroup string
}

// generally it is made by p1:=Person{}

// create an alias
type PersonOptions func(*Person)

func NewName(name string) PersonOptions {
	return func(p *Person) {
		p.Name = name
	}
}

func NewPerson(options ...PersonOptions) *Person {
	temp := &Person{}
	for _, opt := range options {
		opt(temp)
	}
	return temp
}

func main() {
	// p1:=Person{
	// 	Name: "Abbas",
	// }
	// this sets up a p1 with name as abbas and rest as zero values according to data type

	person1 := NewPerson(NewName("Abbas"))
	fmt.Printf("%s\n", person1.Name)
	// this way we made a fractory function acting as a constructor that can accept arguments like a overloaded constructor

	// Immediately Invoked Function Expression
	tempfunc := func(a int, b int) int {
		return a * b
	}(6, 5)

	fmt.Printf("temp func: %d\n", tempfunc)
}
