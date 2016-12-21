package main

import (
	"github.com/zubairhamed/charlotte/td"
	"log"
	"time"
)

func main() {
	// Instantiate HTTP Client to Port 8001
	// client := &http.Client{}

	thing, err := td.LoadFile("td-temperature.json")
	if err != nil {
		panic(err.Error())
	}

	printThing(thing)

	// Every 1 second, send HTTP packet
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		log.Println(thing)
	}
}

func printThing(t *td.Thing) {
	log.Println("@context", t.Context)
	log.Println("@type", t.Type)
	log.Println("name", t.Name)
	log.Println("base", t.Base)

	log.Println("interactions", len(t.Interactions))
	for _, v := range t.Interactions {
		log.Println("- name", v.Name)
	}
}
