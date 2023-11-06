package shake

import "tcp-simulate/message"

const (
	closed      int8 = iota // 关闭
	SYN_SENT                // 同步已发送
	ESTABLISHED             // 建立连接成功
)

type Client interface {
	Recv()
	Send(message.Message)
}

type client struct {
	state int8
}

func (c *client) Recv() {

}

func (c *client) Send(msg message.Message) {

}

func NewClient() Client {
	return &client{
		state: SYN_SENT,
	}
}
