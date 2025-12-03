package main

import "fmt"

type Address struct {
	Block  string
	street string
}

func (a Address) getAddress() (string, string) {
	return a.Block, a.street
}

type Person struct {
	Name    string
	age     int
	Address Address
}

type Address2 struct {
	Block  string
	street string
}

func (a Address2) getAddress() (string, string) {
	fmt.Print("address")
	return a.Block, a.street
}

type Person2 struct {
	Name string
	age  int
	Address
}

func (p Person2) getAddress() (string, string) {
	fmt.Print("person")
	return p.street, p.Block
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

func (c *Person) setAge(agee int) {
	c.age = agee
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

// embedding in structs kindof act like inheritance but not really

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

	pp1 := Person{
		Name: "person1",
		age:  23,
		Address: Address{
			Block:  "dasda",
			street: "faswa",
		},
	}

	pp2 := Person{
		Name: "person2",
		age:  23,
		Address: Address{
			Block:  "dasda",
			street: "faswa",
		},
	}

	pp1.setAge(1)
	pp2.setAge(2)
	fmt.Printf("%d\n", pp1.age)
	fmt.Printf("%d\n", pp2.age)

	// test_person := Person{
	// 	Name: "test_person",
	// 	age:  23,
	// 	Address: Address{
	// 		Block:  "dasda",
	// 		street: "faswa",
	// 	},
	// }
	// test_person.Address.street
	// test_person.Address.Block
	// test_person.Address.getAddress() // in case of compostion (has a)

	// when embedded (is a) acts like inheritance method overriding etc

	test_person2 := Person2{
		Name: "test_person2",
		age:  23,
		Address: Address{
			Block:  "dasda",
			street: "faswa",
		},
	}

	// test_person2.Address
	test_person2.getAddress() // here the method was overridden

}
