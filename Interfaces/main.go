package main

import "fmt"

// whoever implements the interface methods automatically implements the interface

// example of polymorphism using both value and pointer
type Person interface {
	speak()
}

type Human struct {
	Name string
}

func (h *Human) speak() {
	fmt.Printf("This is a human %s\n", h.Name)
}

type Robot struct {
	Name string
}

func (r Robot) speak() {
	fmt.Printf("This is a Robot %s\n", r.Name)
}

type MYINT = int

// compostion in interfaces
type SteetingWheel interface {
	rotate()
	horn()
}

type Engine interface {
	start()
	stop()
}

type Car interface {
	SteetingWheel
	Engine //composition
}

type HondaCar struct {
	CarName string
	Model   int
}

func (h HondaCar) rotate() {

}

func (h HondaCar) horn() {

}

func (h HondaCar) start() {

}

func (h HondaCar) stop() {

}

func main() {

	var mycar Car = HondaCar{
		CarName: "city",
		Model:   2021,
	}
	fmt.Printf("%v\n", mycar)
	// mycar.

	var Humann Person
	var Robott Person
	var Robott2 Person

	Humann = &Human{
		Name: "sda",
	}

	Robott = &Robot{
		Name: "Rorrr",
	}
	Robott2 = Robot{
		Name: "Rorrr2",
	}

	Humann.speak()
	Robott.speak()
	Robott2.speak()

	var v MYINT
	v = 4
	fmt.Printf("%d\n", v)

	var person Person = &Human{
		Name: "Abbas",
	}
	// fmt.Printf("%s\n",person.Name) cannot access, use assertion

	fmt.Printf("%s\n", person.(*Human).Name)

	//here is a better way

	p, valuee := person.(*Human)
	if valuee {
		fmt.Printf("%s\n", p.Name)

	}

}
