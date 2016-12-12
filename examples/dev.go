package main

import (
	"github.com/zubairhamed/charlotte"
	"github.com/zubairhamed/charlotte/connectors/coap"
	"github.com/zubairhamed/charlotte/connectors/http"
)

// This is the All-in-one version of Charlotte (for Development Purposes Only)
func main() {
	// Run Server with Connectors
	go runServer()

	// Runs Thing Registry
	go runRegistry()

	// Runs WebAPI
	go runWebApi()

	// Runs Lambda Service
	go runLambda()

	<-make(chan struct{})
}

func runServer() {
	server := charlotte.NewServer()

	// Register CoAP Connector
	server.RegisterConnector(coap.NewConnector())

	// Register HTTP Connector
	server.RegisterConnector(http.NewConnector())

	if err := server.Start(); err != nil {
		panic(err.Error())
	}
}

func runRegistry() {
	server := charlotte.NewRegistry()

	if err := server.Start(); err != nil {
		panic(err.Error())
	}
}

func runWebApi() {
	server := charlotte.NewWebApiServer()

	if err := server.Start(); err != nil {
		panic(err.Error())
	}
}

func runLambda() {
	server := charlotte.NewWebFunctionsServer()

	if err := server.Start(); err != nil {
		panic(err.Error())
	}
}
