package main

import (
	"fmt"
	"net"
	"os"

	"net-cat/handlers"
)

func main() {
	port := ":8989"
	if len(os.Args) > 2 {
		return
	}
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}
	ln, err := net.Listen("tcp",port)
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
