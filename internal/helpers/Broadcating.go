package helpers

import (
	"fmt"
	"net"
	"time"

	"chat-app/utils"
)

func Broadcasting(message string, sender net.Conn) {
	for client := range utils.Clients {
		if client != sender {
			client.Write([]byte(message))
			client.Write([]byte(fmt.Sprintf("[%s] [%s]: ", time.Now().Format("2006-01-02 15:04:05"), utils.Clients[client])))
		} else {
			client.Write([]byte(fmt.Sprintf("[%s] [%s]: ", time.Now().Format("2006-01-02 15:04:05"), utils.Clients[sender])))
		}
	}
}
