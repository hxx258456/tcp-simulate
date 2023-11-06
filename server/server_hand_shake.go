package server

import (
	"fmt"
	"net"
)

const (
	LISTEN = iota
	SYN_RECV
)

type Server interface {
	Listen()
	HandleClient(net.Conn)
}

type server struct {
	listener net.Listener
	state    int8
}

func (s *server) Listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}
		println("accept in addr: ", conn.LocalAddr().String())
		s.HandleClient(conn)
	}
}

func (s *server) HandleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	recvMsg := buffer[:n]
	fmt.Println("recvMsg:", recvMsg)
}

func NewServer() Server {
	listener, err := net.Listen("tcp", ":8087")
	if err != nil {
		panic(err)
	}
	return &server{
		listener: listener,
		state:    LISTEN,
	}
}
