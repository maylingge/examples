package main

import (
	"net"
	"fmt"
	
)

func main() {
	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to connect the tcp server")
		return
	}

	c.Write([]byte("V2"))

	for {
		b := make([]byte, 1024)
		n, err := c.Read(b)
		if err != nil {
			fmt.Println("Failed to read status")
			return
		}
		fmt.Println(string(b[:n]))
		var message string
		fmt.Scanln(&message)
		c.Write([]byte(message))
	}

}
