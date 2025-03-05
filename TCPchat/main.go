package main

import (
	"fmt"
	"net"
	"os"

	"chat-app/internal/logger"
	"chat-app/internal/server"
	"chat-app/utils"
)

func main() {
	// Initialize the logger || create/open the log file
	logFile, err := logger.Logger()
	if err != nil {
		fmt.Println("Erreur lors de la configuration du fichier de log:", err)
		return
	}
	defer logFile.Close()

	// validate the usage
	port := ":8989"
	if len(os.Args) > 2 {
		logger.ErrorLogger.Println("[USAGE]: ./TCPChat $port")
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}

	// Create a TCP listener on the specified port
	listner, err := net.Listen("tcp", port)
	if err != nil {
		logger.ErrorLogger.Println("error in listening  : ", err)
		fmt.Println("error in listening  : ", err)
		return
	}
	defer listner.Close()
	logger.InfoLogger.Printf("Server runing at port %s\n", port)
	fmt.Printf("Server runing at port %s\n", port)

	// Accept incoming client connections in a loop
	for {
		conn, err := listner.Accept()
		if err != nil {
			logger.ErrorLogger.Println("error in accepting  : ", err)
			fmt.Println("error in accepting  : ", err)
			continue
		}
		// Ensure that the chat is not full before accepting new clients
		utils.MU.Lock()
		if utils.Counter > 4 {
			logger.InfoLogger.Println("Chat is full. Try later...")
			conn.Write([]byte(utils.Yellow + "Chat is full. Try later...\n" + utils.Reset))
			conn.Close()
			utils.MU.Unlock()
			continue
		}
		utils.Counter++
		logger.InfoLogger.Println("new connection accepted", conn.LocalAddr())

		utils.MU.Unlock()
		// Handle the client connection in a separate goroutine
		go server.Server(conn)
	}
}
