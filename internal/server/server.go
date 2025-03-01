package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"chat-app/internal/helpers"
	"chat-app/internal/validators"
	"chat-app/utils"
)

func Server(conn net.Conn) {
	conn.Write([]byte(utils.WelcomeMessage))
	var username string
	var err error

	for {
		username, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading username:", err)
			CloseConnection(conn, "")
			return
		}
		username = strings.TrimSpace(username)

		if !validators.ValidName(username) || !validators.ValidateLength(username) || !validators.SameName(username) {
			conn.Write([]byte("Invalid username...\n"))
			time.Sleep(1 * time.Second)
			conn.Write([]byte("Enter your name again : "))
			continue
		}

		break
	}

	utils.MU.Lock()
	utils.Clients[conn] = username
	utils.MU.Unlock()

	helpers.Broadcasting(utils.Green+fmt.Sprintf("\n%s has joined the chat...\n"+utils.Reset, username), conn)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", err)
			CloseConnection(conn, username)
			return
		}
		message = strings.TrimSpace(message)
		if !validators.ValidMessage(message) || !validators.ValidateLengthMessage(message) {
			time.Sleep(1 * time.Second)
			conn.Write([]byte(fmt.Sprintf("[%s] [%s]:", utils.Time, username)))
			continue
		}
		helpers.Broadcasting(fmt.Sprintf("\n[%s] [%s]: %s\n", utils.Time, username, message), conn)
	}
}

func CloseConnection(conn net.Conn, username string) {
	utils.MU.Lock()
	if username != "" {
		helpers.Broadcasting(utils.Red+fmt.Sprintf("\n%s has left the chat...\n"+utils.Reset, username), conn)
		delete(utils.Clients, conn)
	}
	utils.Counter--
	utils.MU.Unlock()
	conn.Close()
}
