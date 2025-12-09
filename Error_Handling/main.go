package main

import (
	"errors"
	"fmt"
	"time"
)

func divide(a int, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("Error occured \n")
	}
	return a / b, nil
}

func favCreature(x []string) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("An error occured: %s\n", r)
		}
	}()

	fmt.Printf("My fav animal is: %s\n", x[len(x)])
}
func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Error occured: %s\n", r)
		} else {
			fmt.Print("No error")
		}
	}()

	// err := errors.New("Action not allowed")
	// fmt.Printf("%s\n", err)

	ans, err := divide(5, 0)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%d\n", ans)
	}

	originalError := errors.New("Original Error")
	wrappedError := fmt.Errorf("Wrapped error on %w\n", originalError)
	fmt.Printf("%s\n", wrappedError)
	unwrappedError := errors.Unwrap(wrappedError)
	fmt.Printf("Unwrapped Error: %s\n", unwrappedError)

	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}

	favCreature(names)
	time.Sleep(500 * time.Millisecond)
	panic("ERRORRRRRRR AAAAAAAAAAAA")
}
