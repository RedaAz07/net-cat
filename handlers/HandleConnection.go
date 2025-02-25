package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var Client = make(map[net.Conn]string, 10)

func HandleConnection(con net.Conn) {
	var MU sync.Mutex

	defer con.Close()

	welcomeMessage := "Welcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    .       | ' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     -'       --'\n" +
		"inter ur name : "

	con.Write([]byte(welcomeMessage))
	username, err := bufio.NewReader(con).ReadString('\n')

	if err != nil || username == "" {
		fmt.Println(" Invalid username !!!", err)
		return
	}

	username = strings.TrimSpace(username)

	MU.Lock()
	Client[con] = username
	MU.Unlock()

	Broadcast(fmt.Sprintf("👋 A new user has joined: %s \n ", username), con)

	for {
		timee := time.Now().Format("[2006-01-02 15:04:05]")
		con.Write([]byte(fmt.Sprintf("%s [%s]: ", timee, username)))
		data, err := bufio.NewReader(con).ReadString('\n')

		if err != nil {
			break
		}

		data = strings.TrimSpace(data)

		Broadcast(fmt.Sprintf("%s [%s]: %s  \n ", timee, username, data), con)
	}

	MU.Lock()
	delete(Client, con)
	MU.Unlock()

	Broadcast(fmt.Sprintf("🚪 User %s has left the chat.  \n ", username), con)
}
