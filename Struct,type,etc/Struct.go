package main

import "fmt"

type Address struct {
	Block  string
	street string
}

type Person struct {
	Name    string
	age     int
	Address Address
}

// composition
type Engine struct {
	Year int
	CC   int
}

type Car struct {
	Name    string `json:name`
	carType string
	Engine
}

// functions associated with structs
func (c Person) getAge() int {
	return c.age
}

// pass by ref to modify original value
func (c *Person) setBlock(blk string) {
	c.Address.Block = blk
}

type human interface {
	getName() string
	getAge() int
}

type human1 struct {
	name string
	age  int
}

// implementation of interface functions by human1 struct
func (h human1) getName() string {
	return h.name
}

func (h human1) getAge() int {
	return h.age
}

func main() {
	person := Person{
		Name: "Abbas",
		age:  23,
		Address: Address{
			Block:  "dasd",
			street: "dsada",
		},
	}

	fmt.Printf("%d\n", person.getAge())
	person.setBlock("p block")
	fmt.Printf("%s\n", person.Address.Block)

	person2 := &Person{
		Name: "Abbas",
		age:  23,
		Address: Address{
			Block:  "dasd",
			street: "dsada",
		},
	}

	fmt.Printf("%s\n", person2.Name)

	human := human1{
		name: "Abbas",
		age:  23,
	}

	fmt.Printf("%d\n", human.getAge())

	car := Car{
		Name:    "Honda",
		carType: "Sedan",
		Engine: Engine{
			Year: 2021,
			CC:   1500,
		},
	}

	fmt.Printf("%d\n", car.Year)

	adr := &Address{
		Block:  "vfgs",
		street: "gbff",
	}
	fmt.Printf("%s, %s", adr.Block, adr.street)
}
