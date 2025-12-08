package main

import (
	"fmt"
	"sync"
	"time"
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

func send(ch chan int) {
	ch <- 1
}

func send2(ch chan int) {
	ch <- 2
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

	// chl := make(chan<- int)
	// go sendmessage(chl, 2)
	// // tempoo:=<-chl
	// close(chl)

	//temp3 := <-chl // doesnt work due to send only channel

	chl2 := make(chan int)
	go sendmessage2(chl2, 2)
	for v := range chl2 {
		fmt.Printf("%d\n", v)
	}

	// recieve only channel
	// chl3:=make(<-chan int)

	//channel conversions

	// channell:=make(chan int)
	// consider a uni directional channel

	// sendch:=chan<-int (channell) send only
	// recvch:=<-chan int (channell) recv only

	// we have select that is similar to switch but for channels

	channel1 := make(chan int)
	channel2 := make(chan int)

	go send(channel1)
	go send2(channel2)

	select {
	case msg1 := <-channel1:
		{
			fmt.Printf("message: %d\n", msg1)
		}
	case msg2 := <-channel2:
		{
			fmt.Printf("message: %d\n", msg2)
		}
		// default:
		// 	{
		// 		print("none")
		// 	}

	}

	// in case of when multiple channels are ready, it selects one randomly

	// we can make all channels run by adding a loop

	channel3 := make(chan int)
	channel4 := make(chan int)

	go send(channel3)
	go send2(channel4)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel3:
			{
				fmt.Printf("message: %d\n", msg1)
			}
		case msg2 := <-channel4:
			{
				fmt.Printf("message: %d\n", msg2)
			}
		case <-time.After(3 * time.Second):
			{
				print("none")
			}
		}
	}

	// this makes all channels run

	// select {}  in case of empty, it keeps waiting for a case and so is a blocking call

	// producer consumer process
	data := []int{1, 2, 3, 4, 5}
	data_channel := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, val := range data {
			data_channel <- val
		}
		close(data_channel)
	}()

	go func() {
		defer wg.Done()
		for val := range data_channel {
			fmt.Printf("%d\n", val)
		}
	}()

	wg.Wait()

}
