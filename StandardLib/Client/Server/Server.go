package main

import (
	"fmt"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Hello from server\n"))
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Received:", string(buf[:n]))
}

func main() {
	list, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Print("Error listening to port")
		return
	}

	defer list.Close()

	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Print("Error accepting to port")
			return
		}

		go handler(conn)
	}
}
