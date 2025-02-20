package helpers

import (
	"net"
	"sync"
)

var (
	Client = make(map[net.Conn]string, 10)
)

func Broadcast(message string, sender net.Conn) {
 var 	MU     sync.Mutex

	MU.Lock()

	defer MU.Unlock()
	for v := range Client {
		if v != sender {
			v.Write([]byte(message))
		}
	}

}
