package coap

import (
	"github.com/nats-io/go-nats"
	"github.com/zubairhamed/canopus"
	"github.com/zubairhamed/charlotte"
	"log"
	"strings"
)

func NewConnector() charlotte.Connector {
	return &CoapConnector{}
}

type CoapConnector struct {
	charlotte.BaseConnector
	nc *nats.Conn
	server canopus.CoapServer
}

func (c *CoapConnector) Start() (err error) {
	nc, err := charlotte.CreateNatsClient(nats.DefaultURL)
	if err != nil {
		return
	}

	c.nc = nc
	c.server = canopus.NewServer()

	c.server.Get("/td/:device", func(req canopus.Request) canopus.Response {
		msg := req.GetMessage()
		topic := strings.Replace(msg.GetURIPath()[1:], "/", ".", -1)
		log.Println("Publish to", topic)
		if err := c.nc.Publish(topic, msg.GetPayload().GetBytes()); err != nil {
			panic(err.Error())
		}
		
		return canopus.NewResponse(canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment), nil)
	})

	c.server.ListenAndServe(":5683")
	return
}

func (c *CoapConnector) Stop() error {
	return nil
}

