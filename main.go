package main

import (
	"fmt"
	"log"

	"github.com/go-web-dev/event-bus-go-client/client"
)

func main() {
	tcpClient, err := client.New("localhost:8080")
	if err != nil {
		log.Fatal("could not create client", err)
	}
	res, err := tcpClient.HealthOperation()
	if err != nil {
		log.Fatal("could not make health op request", err)
	}
	fmt.Println(res)
}
