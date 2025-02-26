package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"chat-app/internal/helpers"
	// "chat-app/internal/validators"
	"chat-app/utils"
)

func Server(conn net.Conn) {
	conn.Write([]byte(utils.WelcomeMessage))

	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
	username = strings.TrimSpace(username) // bax n7ydo newline awla spaces li flwl ola flkhr

	// if !validators.Isprintable(username) {
	// 	conn.Write([]byte("Only printable characters are allowed.\n"))
	// 	conn.Close()
	// 	return
	// }
	Time := time.Now().Format("2006-01-02 15:04:05")

	utils.MU.Lock()
	utils.Clients[conn] = username
	utils.MU.Unlock()

	helpers.Broadcasting(fmt.Sprintf("\n%s has join the chat... 👋\n", username), conn)
	helpers.Broadcasting((fmt.Sprintf("[%s] [%s]: ", Time, username)),conn)
	for {
		conn.Write([]byte(fmt.Sprintf("[%s] [%s]: ", Time, username)))

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}

		message = strings.TrimSpace(message)

		helpers.Broadcasting(fmt.Sprintf("\n[%s] [%s]: %s\n", Time, username, message), conn)
		helpers.Broadcasting((fmt.Sprintf("[%s] [%s]: ", Time, username)),conn)
	}

	// Remove user when they disconnect
	utils.MU.Lock()
	delete(utils.Clients, conn)
	utils.MU.Unlock()
	helpers.Broadcasting(fmt.Sprintf("\n%s has left the chat... 🚪", username), conn)
	conn.Close()
}
