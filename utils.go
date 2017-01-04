package charlotte

import (
	"github.com/nats-io/go-nats"
	"bytes"
	"encoding/gob"
)

func CreateNatsClient(url string) (nc *nats.Conn, err error) {
	nc, err = nats.Connect(nats.DefaultURL)

	return
}

func CreateMessage(v map[string] string, b []byte) (m *Message) {
	return &Message{
		Vars: v,
		Payload: b,
	}
}

func EncodeMessage(m *Message) []byte {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	enc.Encode(m)

	return buf.Bytes()
}

func DecodeMessage(b []byte) *Message {
	var msg *Message
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	dec.Decode(&msg)

	return msg
}