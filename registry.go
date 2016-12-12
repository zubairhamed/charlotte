package charlotte

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func NewRegistry() *Registry {
	return &Registry{}
}

type Registry struct {
	nc *nats.Conn
	db *bolt.DB
}

func (s *Registry) Start() error {
	nc, err := CreateNatsClient(nats.DefaultURL)
	if err != nil {
		return err
	}

	nc.Subscribe("td.*", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	s.nc = nc

	// TODO Connect to BoltDB
	db, err := bolt.Open("registry.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return nil
}