package charlotte

import "github.com/nats-io/go-nats"

func CreateNatsClient(url string) (nc *nats.Conn, err error){
	nc, err = nats.Connect(nats.DefaultURL)

	return
}
