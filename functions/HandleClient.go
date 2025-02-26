package functions

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	Clients = make(map[net.Conn]string, 10)
	MU      sync.Mutex
)

func HandleClient(conn net.Conn) {
	// defer conn.Close()
	buff := make([]byte, 15)
	// writes data to the connection
	conn.Write([]byte(Welcoming()))
	lenn, err := conn.Read(buff)
	if err != nil {
		fmt.Println(err)
		return
	}

	MU.Lock()
	Clients[conn] = string(buff[:lenn-1])
	MU.Unlock()

	Broadcast("One member has join the chat "+string(Clients[conn]), conn)

	time := time.Now().Format("2006-01-02 15:04:05")

	conn.Write([]byte(fmt.Sprintf("[%s] [%s] : ", time, string(Clients[conn]))))
}
