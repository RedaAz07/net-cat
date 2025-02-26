package functions

import (
	"fmt"
	"net"
	"sync"
	"time"

	"net-cat/helpers"
)

var (
	Clients = make(map[net.Conn]string, 10)
	MU      sync.Mutex
)

func HandleClient(conn net.Conn) {
	// defer conn.Close()
	buff := make([]byte, 1024)
	conn.Write([]byte(Welcoming()))
	lenn, err := conn.Read(buff)
	if err != nil {

		MU.Lock()
		username, exists := Clients[conn]
		if exists {

			Broadcast("user left the chat "+username, conn)
			delete(Clients, conn)
		}
		MU.Unlock()
		return
	}
	for {
		if (lenn > 15 || lenn < 4) || helpers.Invalide(string(buff[:lenn-1])) {
			conn.Write([]byte(" invalide username :"))

			buff = make([]byte, 1024)

			lenn, err = conn.Read(buff)
			if err != nil {
				MU.Lock()
				username, exists := Clients[conn]
				if exists {
					Broadcast("user left the chat "+username, conn)
					delete(Clients, conn)
				}
				MU.Unlock()
				return
			}
		} else {
			break
		}
	}

	MU.Lock()
	Clients[conn] = string(buff[:lenn-1])
	MU.Unlock()
	Broadcast("One member has join the chat "+string(Clients[conn]+"\n"), conn)


	time := time.Now().Format("2006-01-02 15:04:05")

	conn.Write([]byte(fmt.Sprintf("[%s] [%s] :", time, string(Clients[conn]))))
}
