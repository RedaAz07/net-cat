package main

import (
	"fmt"
	"net"
	"os"

	"chat-app/internal/server"
)

func main() {
	port := ":8989"
	if len(os.Args) > 2 {
		port = os.Args[1]
	}
	fmt.Printf("Server runing at port %s\n", port)
	listner, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("error in listening  : ", err)
		return
	}
	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("error in accepting  : ", err)
			return
		}
		go server.Server(conn)
	}
}
