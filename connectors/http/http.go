package http

import (
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"github.com/zubairhamed/charlotte"
	"log"
	"net/http"
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

	r := mux.NewRouter()

	// Get Device
	r.HandleFunc("/td/{dev}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("GET")

	// Update Device
	r.HandleFunc("/td/{dev}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// Get Property
	r.HandleFunc("/td/{dev}/p/{prop}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("GET")

	// Update Property
	r.HandleFunc("/td/{dev}/p/{prop}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// Update Task (Actuate)
	r.HandleFunc("/td/{dev}/a/{action}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// Cancel Task
	r.HandleFunc("/td/{dev}/a/{action}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("DELETE")

	// /td/{dev}/e/{event}
	// POST+LOng Polling
	// Subscribe

	// Update Subscription
	r.HandleFunc("/td/{dev}/e/{event}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// Cancel Subscription
	r.HandleFunc("/td/{dev}/e/{event}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("DELETE")

	go func() {
		http.Handle("/", r)
	}()
	log.Println("HTTP Coinnector Started on Port 8001")

	return
}

func (c *HttpConnector) Stop() error {
	return nil
}
