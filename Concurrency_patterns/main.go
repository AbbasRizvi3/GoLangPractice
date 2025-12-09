package main

import (
	"fmt"
)

// pipelining

// done is made for early stopping the pipline. In case the operation is to be performed onto a partial dataset sent

// other patterns: worker pool, fan in & out
func channelWay(d <-chan struct{}, x ...int) <-chan int {
	channel := make(chan int)
	go func() {
		defer close(channel)
		for _, value := range x {
			select {
			case channel <- value:
			case <-d:
				{
					return
				}
			}
		}
	}()

	return channel
}

func sq(d <-chan struct{}, channel <-chan int) <-chan int {
	channel2 := make(chan int)
	go func() {
		defer close(channel2)
		for value := range channel {
			select {
			case channel2 <- value * 2:
			case <-d:
				{
					return
				}
			}
		}
	}()

	return channel2
}

func main() {
	done := make(chan struct{})
	ch := channelWay(done, 1, 2, 3, 4, 5)
	ch2 := sq(done, ch)
	// time.Sleep(5 * time.Millisecond)
	close(done)
	for value := range ch2 {
		fmt.Printf("%d\n", value)
	}

}
