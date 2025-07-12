# Metasploit RPC Client

## Core Features and Functionallity

- Cmmunicate with Metasploit to list and interact with established Meterpreter sessions.

- Use MessagePack to communicate with Metasploit

## Progress

[X] Establish a connetion to Metasploit
[X] Retreive list of current sessions

## Next Features

[ ] Stop a session
[ ] Run an exploit

## Directory Structure

```
msf-rpc-client
├── bin
│   └── cli
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
└── rpc
    └── msf.go
```
