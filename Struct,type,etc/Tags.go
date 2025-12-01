package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street   string
	House_No int
}

// these tags depend on the library that will be used and are not read by go compiler
type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Blood_Group string `json:"-"`
	Hobby       string `json:"hobby,omitempty"`
	Address     Address
}

// can use inline to spread the address but works in some libs

func main() {
	person := &Person{
		Name:        "Abbas",
		Age:         23,
		Blood_Group: "B+",
		Address: Address{
			Street:   "dasd",
			House_No: 23,
		},
	}

	serialized, _ := json.Marshal(person)
	fmt.Printf("%v\n", string(serialized))
	var person2 Person
	json.Unmarshal(serialized, &person2)
	fmt.Printf("%v\n", person2)

}
