package main

import (
	"fmt"
	"net"
)

func writeResponse(conn net.Conn, body string){
	response := "HTTP/1.1 200 OK\r\n" +
	"Content-Type: text/html\r\n" +
	fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
	"\r\n" +
	body

	conn.Write([]byte(response))
}