package main

func main() {
	ch := make(chan int) // Creating an unbuffered channel

	go func() {
		// No sender is sending a value to the channel
		<-ch // Receiving a value from the channel
	}()

	// The program will deadlock here because there is no sender to send a value
}
