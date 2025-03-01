# Chat App

A simple TCP-based chat application built using Golang. This project allows multiple clients to connect to a server and communicate in real time.

## Features
- Multi-client support
- Usernames for each client
- Welcome message broadcasted to all users when a new user joins
- Message timestamps
- Proper handling of client disconnections

## Installation & Setup

### Prerequisites
- Go (latest version recommended)

### Clone the Repository
```sh
git clone https://github.com/yourusername/chat-app.git
cd chat-app
```

### Run the Server
```sh
go run cmd/server/main.go
```

### Run the Client
Open multiple terminals and run:
```sh
go run cmd/client/main.go
```

## Project Structure
```
chat-app/
├── cmd/
│   ├── server/
│   │   └── main.go
│   ├── client/
│   │   └── main.go
├── internal/
│   ├── helpers/
│   │   ├── broadcast.go
│   ├── validators/
│   │   ├── input.go
├── utils/
│   ├── globals.go
├── README.md
```

## Usage
1. Start the server.
2. Start multiple clients.
3. Enter a username when prompted.
4. Start chatting!

## Example Output
```
Welcome to the chat server!
Enter your username: Alice
Welcome 👋 Alice
[2025-02-26 10:30:00] [Alice]: Hello everyone!
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is open-source and available under the MIT License.

