package main

import (
	"fmt"
	"net/http"

	"github.com/olahol/melody"
)

var m *melody.Melody

func handleWSRequest(w http.ResponseWriter, r *http.Request) {
	m.HandleRequest(w, r)
}

func onConnect(s *melody.Session) {
	fmt.Println("New client connected:", s.Request.RemoteAddr)
	s.Write([]byte("Welcome to the server!"))
}

func onMessage(s *melody.Session, msg []byte) {
	fmt.Println("Message from", s.Request.RemoteAddr, ":", string(msg))
	// Broadcast the message to all clients
	m.Broadcast(msg)
}

func onDisconnect(s *melody.Session) {
	fmt.Println("Client disconnected:", s.Request.RemoteAddr)
}

func onError(s *melody.Session, err error) {
	fmt.Println("Error for client", s.Request.RemoteAddr, ":", err)
}

func main() {
	m = melody.New()

	// register HTTP handler
	http.HandleFunc("/ws", handleWSRequest)
	m.HandleConnect(onConnect)
	m.HandleMessage(onMessage)
	m.HandleDisconnect(onDisconnect)
	m.HandleError(onError)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
