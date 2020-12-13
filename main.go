package main

import (
	"fmt"
	"log"

	"github.com/go-web-dev/event-bus-go/client"
)

func main() {
	eventBusClient, err := client.New("localhost:8080")
	if err != nil {
		log.Fatal("could not create client", err)
	}
	event := map[string]string{
		"k1": "value 1",
		"k2": "value 2",
	}
	res, err := eventBusClient.WriteEvent("hello13", event)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
