package main

import (
	"fmt"
	"bytes"
	"time"
	"net/http"
	"math/rand"
	"log"
	"github.com/zubairhamed/canopus"
)

// Simulates multiple devices using multiple protocols
func main() {
	go startHttpDevice()
	go startCoapDevice()
	<-make(chan struct{})
}

func startHttpDevice() {
	client := &http.Client{} // Instantiate HTTP Client to Port 8001
	ep := "http://localhost:8001/t/temperature-sensor-1/p/temperature"

	// Every 1 second, send HTTP packet
	ticker := time.NewTicker(time.Second * 5)
	rand.Seed(42)

	for range ticker.C {
		val := fmt.Sprintf("%d", rand.Intn(50))
		log.Println("[HTTP] Sending out", val)
		r := bytes.NewReader([]byte(val))


		req, _ := http.NewRequest("PUT", ep, r)
		req.Header.Set("Content-Type", "application/json")

		_, err := client.Do(req)
		if err != nil {
			log.Println("[HTTP]", err.Error())
		}
	}
}

func startCoapDevice() {
	conn, err := canopus.Dial("localhost:5683")
	if err != nil {
		panic(err.Error())
	}

	// Every 1 second, send HTTP packet
	ticker := time.NewTicker(time.Second * 5)
	rand.Seed(42)

	for range ticker.C {
		val := fmt.Sprintf("%d", rand.Intn(50))
		log.Println("[CoAP] Sending out", val)

		req := canopus.NewRequest(canopus.MessageConfirmable, canopus.Get).(*canopus.CoapRequest)
		req.SetStringPayload(val)
		req.SetRequestURI("/t/temperature-sensor-2/p/temperature")

		_, err := conn.Send(req)
		if err != nil {
			log.Println("[CoAP]", err.Error())
		}
	}
}
