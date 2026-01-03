package main

import (
	"fmt"
	"net"
)

func main(){
	ln, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening on port: 8080")

	for{
		conn, err := ln.Accept()
		if err != nil{
			continue
		}
		go handleConnection(conn)
	}
}

