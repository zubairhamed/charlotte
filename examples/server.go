package main

import (
	"github.com/zubairhamed/charlotte"
	"github.com/zubairhamed/charlotte/connectors/coap"
	"github.com/zubairhamed/charlotte/connectors/http"
	"github.com/zubairhamed/charlotte/webapis/w3c"
)

// This is the All-in-one version of Charlotte (for Development Purposes Only)
func main() {
	// Run Server with Connectors
	go runServer()

	// Runs Thing Registry
	go runThingsRegistry()

	go runCredsRegistry()

	// Runs WebAPI
	go runWebApi()

	// Runs Lambda Service
	// go runLambda()

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

func runCredsRegistry() {
	server := charlotte.NewCredsRegistry()

	if err := server.Start(); err != nil {
		panic(err.Error())
	}
}

func runThingsRegistry() {
	server := charlotte.NewThingRegistry()

	if err := server.Start(); err != nil {
		panic(err.Error())
	}
}

func runWebApi() {
	server := charlotte.NewWebApiServer()

	server.RegisterWebInterface(wc3.NewWebInterface())

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
