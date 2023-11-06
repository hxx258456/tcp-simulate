package client

import (
	"net"
	"tcp-simulate/message"
)

const (
	SYN_SENT = iota
)

type Client interface {
	Recv()
	Send(message.Message)
}

type client struct {
	state int8
	conn  net.Conn
}

func (c *client) Recv() {

}

func (c *client) Send(msg message.Message) {
}

func NewClient() Client {
	conn, err := net.Dial("tcp", "127.0.0.1:8087")
	if err != nil {
		panic(err)
	}
	return &client{
		state: SYN_SENT,
		conn:  conn,
	}
}
