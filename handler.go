package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Println("New client connected:",conn.RemoteAddr())

	//Read the request line 
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	method := parts[0]
	path := parts[1]

	fmt.Println("Method:", method)
	fmt.Println("Path:", path)

	//consume remaining headers
	for {
		line,_ := reader.ReadString('\n')
		if (line == "\r\n"){
			break
		}
	}

	//routing
	var body string
	if path == "/" {
		body = "<h1>Home Page</h1>"
	} else if path == "/contact" {
		body = "<h1>Contact Page</h1>"
	} else {
		body = "<h1>404 Not Found</h1>"
	}

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
		"\r\n" +
		body

	conn.Write([]byte(response))
	
}