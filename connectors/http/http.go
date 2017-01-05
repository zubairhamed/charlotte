package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"github.com/zubairhamed/charlotte"
	"io/ioutil"
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
		log.Println(err)
		return
	}
	c.nc = nc

	r := mux.NewRouter()
	c.initRoutes(r)

	http.ListenAndServe(":8001", r)
	return
}

func (c *HttpConnector) initRoutes(r *mux.Router) {
	r.HandleFunc("/t/{dev}", c.onGetThing).Methods("GET")
	r.HandleFunc("/t/{dev}/model", c.onGetThingModel).Methods("GET")
	r.HandleFunc("/t/{dev}", c.onUpdateThing).Methods("PUT")
	r.HandleFunc("/t/{dev}/p/{prop}", c.onGetProperty).Methods("GET")
	r.HandleFunc("/t/{dev}/p/{prop}", c.onUpdateProperty).Methods("PUT")
	r.HandleFunc("/t/{dev}/a/{action}", c.onUpdateTask).Methods("PUT")
	r.HandleFunc("/t/{dev}/a/{action}", c.onCancelTask).Methods("DELETE")
	r.HandleFunc("/t/{dev}/e/{event}", c.onUpdateSubscription).Methods("PUT")
	r.HandleFunc("/t/{dev}/e/{event}", c.onCancelSubscription).Methods("DELETE")
}

func (c *HttpConnector) onUpdateProperty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dev := vars["dev"]
	prop := vars["prop"]

	t := fmt.Sprintf("t.%s.p.%s", dev, prop)

	b, _ := ioutil.ReadAll(r.Body)

	c.nc.Publish(t, charlotte.EncodeMessage(charlotte.CreateMessage(vars, b)))
}

func (c *HttpConnector) onGetThing(w http.ResponseWriter, r *http.Request)           {}
func (c *HttpConnector) onGetThingModel(w http.ResponseWriter, r *http.Request)      {}
func (c *HttpConnector) onUpdateThing(w http.ResponseWriter, r *http.Request)        {}
func (c *HttpConnector) onGetProperty(w http.ResponseWriter, r *http.Request)        {}
func (c *HttpConnector) onUpdateTask(w http.ResponseWriter, r *http.Request)         {}
func (c *HttpConnector) onCancelTask(w http.ResponseWriter, r *http.Request)         {}
func (c *HttpConnector) onUpdateSubscription(w http.ResponseWriter, r *http.Request) {}
func (c *HttpConnector) onCancelSubscription(w http.ResponseWriter, r *http.Request) {}

func (c *HttpConnector) Stop() error {
	return nil
}
