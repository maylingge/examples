package main

import (
	"net"
	"fmt"
	"crypto/tls"
	"github.com/mytry/tcpwithtls/common"
)

func handleConnection(c net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buf[:n]))
		c.Write(buf[:n])

	}
}

func main() {
	config := common.MustGetTlsConfiguration()
	ln, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		fmt.Println("Failed to create tcp server on :8080!")
		return
	}

	for {
		fmt.Println("Waiting for incoming connection...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connect")
			return
		}
		fmt.Printf("Connection from: %s \n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
