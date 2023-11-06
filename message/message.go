package message

type Message struct {
	Header Header
	Body   []byte
}

type Header struct {
	ACK  int
	SEQ  int
	FLAG FLAG
}
