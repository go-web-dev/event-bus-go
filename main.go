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
	res, err := eventBusClient.MarkEvent("4f8c8a39-3b7b-4316-bfb2-d0c6ad9f0ccb", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
