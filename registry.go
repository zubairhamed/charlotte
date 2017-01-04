package charlotte

import (
	"github.com/boltdb/bolt"
	"github.com/nats-io/go-nats"
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/zubairhamed/charlotte/td"
)

func NewRegistry() *Registry {
	return &Registry{
		db: &BoltDB{},
	}
}

type Registry struct {
	nc *nats.Conn
	db *BoltDB
}

func (s *Registry) onUpdateThingProperty(m *nats.Msg) {
	// Convert to Message
	msg := DecodeMessage(m.Data)

	vars := msg.Vars
	devId := vars["dev"]

	s.db.UpdateProperty(devId, vars["prop"], string(msg.Payload))
}

func (s *Registry) Start() error {
	log.Println("Starting Thing Registry")
	nc, err := CreateNatsClient(nats.DefaultURL)
	if err != nil {
		return err
	}

	_, err = nc.Subscribe("t.*.p.>", s.onUpdateThingProperty)

	if err != nil {
		panic(err.Error())
		return err
	}

	s.nc = nc

	//return s.db.Setup()
	return nil
}

type BoltDB struct {

}

func (b *BoltDB) Setup() error {
	db, err := b.open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("td"))

		content, _ := ioutil.ReadFile("./td-temperature.json")
		b.Put([]byte("temperature-sensor-1"), content)

		b, _ = tx.CreateBucket([]byte("twin"))

		content, _ = ioutil.ReadFile("./td-temperature-instance.json")
		b.Put([]byte("temperature-sensor-1"), content)

		b, _ = tx.CreateBucket([]byte("syslogs"))
		return nil
	})
	return e
}

func (b *BoltDB) open() (*bolt.DB, error) {
	return bolt.Open("registry.db", 0600, nil)
}

func (b *BoltDB) close(db *bolt.DB) {
	db.Close()
}

func (b *BoltDB) GetThing(id string) map[string]interface{} {
	return make(map[string]interface{})
}

func (b *BoltDB) UpdateProperty(id, prop, value string) error {
	db, err := b.open()
	if err != nil {
		return err
	}
	defer b.close(db)

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("twin"))

		v := b.Get([]byte(id))
		var m map[string]interface{}

		json.Unmarshal(v, &m)

		m[prop] = value

		content, err := json.Marshal(m)
		if err != nil {
			return err
		}

		b.Put([]byte(id), content)

		return nil
	})
}

// TODO
func (b *BoltDB) GetThingModel() (thing *td.Thing, err error) {
	db, err := b.open()
	if err != nil {
		return
	}
	defer db.Close()

	return
}

func (b *BoltDB) DeleteThing() (err error){
	db, err := b.open()
	if err != nil {
		return err
	}
	defer db.Close()

	return
}

func (b *BoltDB) UpdateThing() (err error) {
	db, err := b.open()
	if err != nil {
		return err
	}
	defer db.Close()

	return
}
