package main

import (
	"fmt"
	"log"

	"github.com/go-web-dev/event-bus-go/client"
)

func main() {
	credentials := client.Credentials{
		ClientID: "dope_go_client_id",
		ClientSecret: "dope_go_client_secret",
	}
	eventBusClient, err := client.New("localhost:8080", credentials)
	if err != nil {
		log.Fatal("could not create client", err)
	}
	res, err := eventBusClient.CreateStream("steve")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
