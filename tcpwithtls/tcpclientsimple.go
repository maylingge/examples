package main

import (
	"fmt"
	"crypto/tls"
	"github.com/mytry/tcpwithtls/common"
)

func main() {
	config := common.MustGetTlsConfiguration()
	config.ServerName = "localhost"
	c, err := tls.Dial("tcp", ":8080", config)
	if err != nil {
		fmt.Println("failed to connect the tcp server")
		fmt.Println(err)	
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
