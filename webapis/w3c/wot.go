package wc3

import (
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"github.com/zubairhamed/charlotte"
	"log"
	"net/http"
)

func NewWebInterface() charlotte.WebInterface {
	return &WotWebInterface{}
}

type WotWebInterface struct {
	charlotte.BaseConnector
	nc *nats.Conn
}

func (c *WotWebInterface) Start() (err error) {
	nc, err := charlotte.CreateNatsClient(nats.DefaultURL)
	if err != nil {
		return
	}

	c.nc = nc

	// TODO: Start HTTP Server on Port 8082?

	/*
		/td/{dev}
		GET
		Get Thing

		/td/{dev}
		PUT
		Update Thing

		/td/{dev}/p/{prop}
		GET
		Get Property


	*/

	// :8082/td/deviceId
	r := mux.NewRouter()
	r.HandleFunc("/td/{id}", func(w http.ResponseWriter, r *http.Request) {

	})

	log.Println("WoT Web Interface Started on Port 8082")
	http.ListenAndServe(":8082", r)
	return
}

func (c *WotWebInterface) Stop() error {
	return nil
}
