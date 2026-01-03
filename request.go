package main

import (
	"bufio"
	"strings"
)

type Request struct {
	Method string
	Path string
}

func parseRequest(reader *bufio.Reader) Request {
	line, error := reader.ReadString('\n')
	if error != nil{
		panic(error)
	}
	parts := strings.Split(strings.TrimSpace(line)," ")

	req := Request{
		Method:parts[0],
		Path: parts[1],
	}

	for{
		line,_ := reader.ReadString('\n')
		if (line =="\r\n") {
			break
		}
	}

	return req
}