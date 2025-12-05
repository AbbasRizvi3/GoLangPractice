package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Person struct {
	Name string
	Age  int
}

func sendMessage[T int | Person](chl chan T, data T) {
	chl <- data
}

func sendmessage(chl chan<- int, data int) {
	chl <- data
}

func sendmessage2[T int](chl chan T, data T) {
	chl <- data
	chl <- data * 2
	chl <- data * 3
	close(chl)
}

func main() {

	// unbuffered channels are blocking in nature, both sender and reciever block so no need to use wait or done
	// in case of buffered, it is unblocking until the buffer is filled
	// we might use wait etc in case multiple senders are there for a certain channel
	ic := make(chan int)
	ic2 := make(chan Person)
	go sendMessage(ic, 6)
	temp := <-ic
	close(ic)
	fmt.Printf("%d\n", temp)

	person := Person{
		Name: "Abbas",
		Age:  23,
	}
	go sendMessage(ic2, person)
	temp2 := <-ic2
	close(ic2)
	fmt.Printf("%v\n", temp2)

	// send only channel

	chl := make(chan<- int)
	go sendmessage(chl, 2)
	close(chl)

	//temp3 := <-chl // doesnt work due to send only channel

	chl2 := make(chan int)
	go sendmessage2(chl2, 2)
	for v := range chl2 {
		fmt.Printf("%d\n", v)
	}

}
