package coap

import (
	"github.com/nats-io/go-nats"
	"github.com/zubairhamed/canopus"
	"github.com/zubairhamed/charlotte"
	"log"
	"strings"
)

func NewConnector() charlotte.Connector {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err.Error())
	}

	return &CoapConnector{
		BaseConnector: charlotte.BaseConnector{
			NatsClient: nc,
		},
	}
}

type CoapConnector struct {
	charlotte.BaseConnector
	server canopus.CoapServer
}

func (c *CoapConnector) Start() (err error) {
	c.server = canopus.NewServer()

	c.server.Get("/td/:device", func(req canopus.Request) canopus.Response {
		// TODO: Map Headers
		// TODO: Send to NATS
		msg := req.GetMessage()
		topic := strings.Replace(msg.GetURIPath()[1:], "/", ".", -1)
		log.Println("Publish to", topic)
		if err := c.NatsClient.Publish(topic, msg.GetPayload().GetBytes()); err != nil {
			panic(err.Error())
		}

		log.Println("Hello Called", req.GetAttribute("device"))
		msg = canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		// msg.SetStringPayload("Acknowledged with response : " + req.GetMessage().GetPayload().String())

		res := canopus.NewResponse(msg, nil)
		return res
	})

	c.server.ListenAndServe(":5683")
	return
}

func (c *CoapConnector) Stop() error {
	return nil
}

/*
server := canopus.NewServer()

	server.Get("/hello", func(req canopus.Request) canopus.Response {
		log.Println("Hello Called")
		msg := canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		msg.SetStringPayload("Acknowledged with response : " + req.GetMessage().GetPayload().String())

		res := canopus.NewResponse(msg, nil)
		return res
	})

	server.Post("/hello", func(req canopus.Request) canopus.Response {
		log.Println("Hello Called via POST")

		msg := canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		msg.SetStringPayload("Acknowledged: " + req.GetMessage().GetPayload().String())
		res := canopus.NewResponse(msg, nil)
		return res
	})

	server.Get("/basic", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewPlainTextPayload("Acknowledged"))
		res := canopus.NewResponse(msg, nil)
		return res
	})

	server.Get("/basic/json", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), nil)

		res := canopus.NewResponse(msg, nil)

		return res
	})

	server.Get("/basic/xml", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), nil)
		res := canopus.NewResponse(msg, nil)

		return res
	})

	server.OnMessage(func(msg canopus.Message, inbound bool) {
		canopus.PrintMessage(msg)
	})

	server.ListenAndServe(":5683")
	<-make(chan struct{})
*/
