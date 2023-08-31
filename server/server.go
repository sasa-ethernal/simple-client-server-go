package main

import (
	"fmt"
	"net"
	"os"
	"messaging-test/types"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	msgBytes := make([]byte, 1024)
	n, err := conn.Read(msgBytes)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	message := string(msgBytes[:n])

	msg, err := types.DecodeMessage(msgBytes[:n])
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}
	fmt.Println("Received message from:", conn.RemoteAddr(), message)

	//Forward the message to Client 2
	client2Conn, err := net.Dial("tcp", msg.Destination)
	if err != nil {
		fmt.Println("Error connecting to Client 2:", err)
		return
	}
	defer client2Conn.Close()

	_, err = client2Conn.Write(msgBytes[:n])
	if err != nil {
		fmt.Println("Error forwarding message to Client 2:", err)
		return
	}
	fmt.Println("Message forwarded to:", msg.Destination, message)
}

func main() {
	// serverAddr := "localhost:18081"
	serverAddr := os.Args[1]

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on", serverAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
