# Metasploit RPC Client

## Core Features and Functionality

- Communicate with Metasploit to list and interact with established Meterpreter sessions.
- Use MessagePack to communicate with Metasploit.

## Progress

- [x] Establish a connection to Metasploit
- [x] Retrieve list of current sessions
- [x] Implement reverse proxy for routing to Metasploit listeners
- [x] Stop a session

## Next Features

- [ ] Run an exploit

## Configuration

To run this tool, you need to create a `.msf-skynet.yml` file in your home directory
(`~/.msf-skynet.yml`).
This configuration file should contain the necessary settings to connect to the Metasploit
RPC server, such as the host, port, username, and password.
Below is an example configuration:

```sh
msfhost: "localhost:55552"
msfpass: "secret"
msfuser: "msf"
```

## Installation and Running

### Option 1: Install via Go

To install the Metasploit RPC Client directly using Go, run the following command:

```sh
go install github.com/dracuxan/msf-skynet@latest
```

This will download and install the binary to your $GOPATH/bin directory.
Ensure that $GOPATH/bin is in your system's PATH to run the msf-skynet command from anywhere.
To run the client after installation:

```sh
msf-skynet
```

Option 2: Set Up Locally
To set up and run the project locally, follow these steps:

Clone the repository:

```sh
git clone https://github.com/dracuxan/msf-skynet.git && cd msf-skynet
```

Install dependencies:Ensure you have Go installed, then run:

```sh
go mod tidy
```

Build and run:Build the project using:

```sh
go build -o msf-skynet
```

Then run the binary:

```sh
./msf-skynet
```

Alternatively, use the provided Makefile (if available) to build and run:

```sh
make build
make run
```

## Available Commands:

1. List Sessions

```sh
msf-skynet sessionList
```

2. Run Multiplexer

```sh
msf-skynet multiplexer [--port <port>] [--mapping <hostname=url>]
```

Starts a reverse proxy server that routes HTTP requests to Metasploit listeners based on the Host header.

Flags:

- `--port <int>`: Port for the reverse proxy to listen on (default: 80). Requires privileged access for ports below 1024.

- `--mapping <string>`: Hostname to URL mappings (e.g., attacker1.com=http://10.0.1.20:10080). Can be specified multiple times for multiple mappings.

Example:

```sh
msf-skynet multiplexer --port 8080 --mapping attacker1.com=http://10.0.1.20:10080 --mapping attacker2.com=http://10.0.1.20:20080
```

## Directory Structure

```sh
msf-skynet
├── msf
│ ├── auth.go
│ ├── helper.go
│ ├── models.go
│ ├── multiplexer.go
│ └── sessions.go
│
├── handlers
│   └── getSessions.go
│
├── cmd
│ ├── root.go
│ ├── multiplexer.go
│ └── sessionList.go
│
├── main.go
│
├── go.mod
├── go.sum
│
├── Makefile
├── LICENSE
└── README.md
```
