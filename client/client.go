package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// serverAddr := "localhost:18081"
	hostAddr := os.Args[1]
	serverAddr := os.Args[2]

	listener, err := net.Listen("tcp", hostAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on", hostAddr)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go handleConnection(conn)
		}
	}()

	// fmt.Println("Message to send:")
	for {
		var msg string
		_, err := fmt.Scan(&msg)

		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			os.Exit(1)
		}
		defer conn.Close()

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	
		fmt.Println("Message sent to server:", msg)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received message:", string(buffer[:]))
}
