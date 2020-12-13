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
	res, err := eventBusClient.GetStreamInfo("hello13")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Body)
}
