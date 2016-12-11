package main

import (
	"github.com/zubairhamed/charlotte"
	"github.com/zubairhamed/charlotte/connectors/coap"
)

func main() {

	server := charlotte.NewServer()

	server.RegisterConnector(coap.NewConnector())

	if err := server.Start(); err != nil {
		panic(err.Error())
	}

	<-make(chan struct{})
}
