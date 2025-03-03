# Net-Cat

A simple TCP-based chat application built using Golang. This project allows multiple clients to connect to a server and communicate in real time.

## Features
- Multi-client support
- Usernames for each client
- Welcome message broadcasted to all users when a new user joins
- Message timestamps
- Proper handling of client disconnections
- Server logging system

## Installation & Setup

### Prerequisites
- Go (latest version recommended)

### Clone the Repository
```sh
git clone https://github.com/ranniz/net-cat.git
cd net-cat
```

### Run the Server
```sh
go run TCPchat/main.go
```
### join the chat  
```sh
nc localhost port
```

## Project Structure
```
NET-CAT/
├── internal/
│   ├── helpers/
│   │   ├── Broadcating.go
│   ├── logger/
│   │   ├── logger.go
│   ├── server/
│   │   ├── server.go
│   ├── validators/
│   │   ├── validators.go
├── logs/
│   ├── server.log
├── TCPchat/
│   ├── main.go
├── utils/
│   ├── utils.go
├── .gitignore
├── go.mod
├── network.md
├── README.md
```

## Usage
1. Start the server.
2. Start multiple clients.
3. Enter a username when prompted.
4. Start chatting!

## Example Output
```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    .       | ' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     -'       --'
[ENTER YOUR NAME]:ahmed
[2025-03-02 11:41:28] [ahmed]: 
```
```
ahmed has joined the chat...
[2025-03-02 11:41:18] [ahmed]: hello
[2025-03-02 11:41:18] [ahmed]: salam
reda has joined the chat...
[2025-03-02 11:41:18] [reda]: bikhir
[2025-03-02 11:41:18] [reda]: hanya
[2025-03-02 11:41:18] [ahmed]: lahfdek
[2025-03-02 11:42:28] [aymen]: 
reda has left the chat...
[2025-03-02 11:42:30] [aymen]: 
ahmed has left the chat...
[2025-03-02 11:42:31] [aymen]: 
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is open-source and available under the MIT License.

