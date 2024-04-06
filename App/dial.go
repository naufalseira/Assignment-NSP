package main

import (
	"fmt"
	"net"
	"time"

	"main.mod/Handler"
)

func main() {
	dial, err := net.DialTimeout("tcp", "localhost:1500", time.Second*5)
	handler.ErrorHandler(err)
	defer dial.Close()

	dial.SetDeadline(time.Now().Add(time.Second * 5))

	netErr, ok := err.(net.Error)

	_, err = dial.Write([]byte("HELLO ABANGKUHH"))
	handler.ErrorHandler(err)

	buffer := make([]byte, 1024)
	n, err := dial.Read(buffer)

	fmt.Println("Server: ", string(buffer[:n]))

	if err != nil {
		if ok && netErr.Timeout() {
			fmt.Println("Connection Timeout!")
			return
		} else {
			handler.ErrorHandler(err)
		}
	}
}
