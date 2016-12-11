package charlotte

import (
	"github.com/nats-io/go-nats"
)

type BaseConnector struct {
	NatsClient *nats.Conn
	abc        string
}
