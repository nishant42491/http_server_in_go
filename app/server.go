package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(c net.Conn) {
	// Read data from the connection

	// Write data to the connection
	// Close the connection

	defer c.Close()

	recievedData := make([]byte, 4096)
	_, err := c.Read(recievedData)
	if err != nil {
		fmt.Println("Error reading data: ", err.Error())
		return
	}

	fmt.Println("Data recieved: ", string(recievedData))

	httpresponseheader := "HTTP/1.1 200 OK\r\n \r\n"

	//write the response header
	_, err = c.Write([]byte(httpresponseheader))

	if err != nil {
		fmt.Println("Error writing data: ", err.Error())
		return
	}

	//print made connection
	fmt.Println("Connection made")

}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	//Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	//handle connection

	fmt.Println("Connection accepted")

	handleConnection(conn)

}
