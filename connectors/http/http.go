package http

import (
	"github.com/nats-io/go-nats"
	"github.com/zubairhamed/charlotte"
)

func NewConnector() charlotte.Connector {
	return &HttpConnector{}
}

type HttpConnector struct {
	charlotte.BaseConnector
	nc *nats.Conn
}

func (c *HttpConnector) Start() (err error) {
	nc, err := charlotte.CreateNatsClient(nats.DefaultURL)
	if err != nil {
		return
	}

	c.nc = nc

	// TODO: Start HTTP Server on Port 8081

	return
}

func (c *HttpConnector) Stop() error {
	return nil
}

