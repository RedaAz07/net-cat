package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"net-cat/functions"
)

func main() {
	// when a client does this :  nc $IP $port ->our server should :
	// listen by default on port '8989'
	// sends him a welcome msg by 'default'
	// Ping to other chat members if someone left or joined
	// uploade old msgs to the user if he joins
	// keep the connection open even if no user is connected
	// save the msgs history in a file

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	port :=  "8989" 
	if len(os.Args) == 2  {
		 port  =  os.Args[1]
	}


	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept() // this accept data that u can manipulate
		if err != nil {
			log.Println(err)
			continue // keep trying to connect
		}
		go functions.HandleClient(conn)
	}
}
