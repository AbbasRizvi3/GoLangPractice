package main

import "fmt"

type circle struct {
	Radius float32
}

func (c circle) GetArea() float32 {
	return c.Radius * c.Radius
}

func main() {
	circle1 := circle{Radius: 4.0}

	fmt.Printf("%f\n", circle1.Radius)
	fmt.Printf("%f\n", circle1.GetArea())

}
