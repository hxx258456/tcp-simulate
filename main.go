package main

import "tcp-simulate/server"

func main() {
	server := server.NewServer()
	server.Listen()
}
