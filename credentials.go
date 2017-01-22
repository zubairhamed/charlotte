package charlotte

import (
	"log"
	"github.com/nats-io/go-nats"
	"github.com/boltdb/bolt"
)

func NewCredsRegistry() *CredentialsRegistry {
	return &CredentialsRegistry{
		db: &CredsBoltDB{},
	}
}

type CredentialsRegistry struct {
	nc *nats.Conn
	db *CredsBoltDB
}

func (s *CredentialsRegistry) Start() (err error) {
	log.Println("Starting Credentials Registry")

	nc, err := CreateNatsClient(nats.DefaultURL)
	if err != nil {
		return err
	}
	s.nc = nc


	// TODO
	/*
		Authenticate
		Add
		Get
		List
		Update
	 */
	return s.db.Setup()
}


type CredsBoltDB struct {

}

func (b *CredsBoltDB) Setup() error {
	return nil
}

func (b *CredsBoltDB) open() (*bolt.DB, error) {
	return bolt.Open("registry.db", 0600, nil)
}

func (b *CredsBoltDB) close(db *bolt.DB) {
	db.Close()
}
