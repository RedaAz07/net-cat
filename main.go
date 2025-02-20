package main

import (
	"fmt"
	"net"
	"net-cat/handlers"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error in listening  : ", err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error in accepting  : ", err)
			return
		}
		go handlers.HandleConnection(conn)
	}
}

