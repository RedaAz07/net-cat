package utils

import (
	"net"
	"sync"
	"time"
)

var (
	Time    = time.Now().Format("2006-01-02 15:04:05")
	Counter = 1
	MU      sync.Mutex
	Clients = make(map[net.Conn]string)
	History []string

	WelcomeMessage = "Welcome to TCP-Chat!\n" +
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
		"[ENTER YOUR NAME]:"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)
