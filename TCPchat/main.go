package main

import (
	"fmt"
	"net"
	"os"

	"chat-app/internal/server"
	"chat-app/utils"
)

func main() {
	port := ":8989"
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
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
		utils.MU.Lock()
		if utils.Counter > 10 {
			conn.Write([]byte("Chat is full. Try later...\n"))
			conn.Close()
			utils.MU.Unlock()
			continue
		}
		utils.Counter++
		utils.MU.Unlock()

		go server.Server(conn)
	}
}
