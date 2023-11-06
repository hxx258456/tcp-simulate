package client

import (
	"net"
	"tcp-simulate/message"
	"testing"
)

func Test_client_Send(t *testing.T) {
	type fields struct {
		state int8
		conn  net.Conn
	}
	type args struct {
		msg message.Message
	}
	conn, err := net.Dial("tcp", "127.0.0.1:8087")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "client Send能力测试",
			fields: fields{
				state: SYN_SENT,
				conn:  conn,
			},
			args: args{
				msg: message.Message{},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				state: tt.fields.state,
				conn:  tt.fields.conn,
			}
			c.Send(tt.args.msg)
		})
	}
}
