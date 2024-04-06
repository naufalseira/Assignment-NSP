package main

import (
	"fmt"
	"net"

	"main.mod/Handler"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1500")
	handler.ErrorHandler(err)

	fmt.Println("-----------------------------")
	fmt.Printf("| Bound to %q |\n", listener.Addr())
	fmt.Println("-----------------------------")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handler.ErrorHandler(err)

		go dialHandler(conn)
	}
}

func dialHandler(conn net.Conn) {
	defer conn.Close()

	fmt.Println("----------------------------")
	fmt.Println("| One Connection Received! |")
	fmt.Println("----------------------------")

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	handler.ErrorHandler(err)

	fmt.Println("Client: ", string(buffer[:n]))

	_, err = conn.Write([]byte("MENYALAAA!!!"))
	handler.ErrorHandler(err)
}
