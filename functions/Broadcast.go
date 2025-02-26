package functions

import (
	"net"
)

func Broadcast(message string, sender net.Conn) {
	for conn := range Clients {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			if err != nil {
				return
			}
		}
	}
}
