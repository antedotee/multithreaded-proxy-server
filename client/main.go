package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Println("failed to create the connection", err)
		return
	}
	defer connection.Close()

	// Send a valid HTTP GET request
	httpRequest := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
	_, err = connection.Write([]byte(httpRequest))
	if err != nil {
		log.Println("failed to send data", err)
		return
	}

	// Read and print the response
	buf := make([]byte, 4096)
	n, err := connection.Read(buf)
	if err != nil {
		log.Println("failed to read response", err)
		return
	}
	fmt.Println(string(buf[:n]))
}