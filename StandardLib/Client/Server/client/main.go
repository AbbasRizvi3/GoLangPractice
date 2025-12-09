package main

import (
	"fmt"
	"net"
)

func main() {
	list, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Print("Error dialing")
	}
	defer list.Close()

	// Read server greeting
	buf := make([]byte, 1024)
	n, _ := list.Read(buf)
	fmt.Println("Server says:", string(buf[:n]))

	// Send message
	list.Write([]byte("Hello server!\n"))
}
