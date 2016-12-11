package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err.Error())
	}

	nc.Subscribe("td.*", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	fmt.Println("Listening..")

	<-make(chan struct{})
}
