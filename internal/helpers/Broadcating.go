package helpers

import (
	"net"

	"chat-app/utils"
)

func Broadcasting(message string, sender net.Conn) {
	for client := range utils.Clients {
		if client != sender {
			client.Write([]byte(message))
		}
	}
}
