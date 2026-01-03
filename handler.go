package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn){
	defer conn.Close()

	fmt.Println("New CLient Connected: ", conn.RemoteAddr())

	reader:= bufio.NewReader(conn)

	req := parseRequest(reader)
	res := handleRoute(req)

	writeResponse(conn, res)
}