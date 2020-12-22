# event-bus-go

[![Build Status](https://travis-ci.org/go-web-dev/event-bus-go.svg?branch=master)](https://travis-ci.org/go-web-dev/event-bus-go)
[![codecov](https://codecov.io/gh/go-web-dev/event-bus-go/branch/master/graph/badge.svg)](https://codecov.io/gh/go-web-dev/event-bus-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-web-dev/event-bus-go)](https://goreportcard.com/report/github.com/go-web-dev/event-bus-go)
[![GoDoc](https://godoc.org/github.com/go-web-dev/event-bus-go/client?status.svg)](https://godoc.org/github.com/go-web-dev/event-bus-go/client)

## Event Bus Go Client library

This project represents the library that communicates between
any Go application aka the client, and the [Event Bus](https://github.com/go-web-dev/event-bus)
micro service aka the server.

The library is responsible for communicating over TCP with
the Event Bus micro service according to a specific protocol.

So in other words, instead of creating connections, closing them
writing specific bytes to the connection (according to the protocol)
so that the Event Bus service understand what the Go application wants,
we instead delegate this responsibility to this client library.

Library Version: `v1.0.3` -  latest stable version

### Operations

The library supports communication for all the following operations:

- `health`
- `create_stream`
- `delete_stream`
- `get_stream_info`
- `get_stream_events`
- `write_event`
- `process_events`
- `retry_events`
- `mark_event`
- `exit`

**Note**: Before proceeding with library usage, make sure to
create a new client with the correct credentials
found inside [config.yaml](https://github.com/go-web-dev/event-bus/blob/master/config/config.yaml)
inside the Event Bus project.

Here's a small example:

```go
import (
    "log"

    "github.com/go-web-dev/event-bus-go/client"
)


credentials := client.Credentials{
    ClientID: "dope_go_client_id",
    ClientSecret: "dope_go_client_secret",
}

eventBusClient, err := client.New("localhost:8080", credentials)
if err != nil {
    log.Fatal("could not create client", err)
}
```

Make sure to adjust `ClientID` and `ClientSecret` accordingly.

### Docs

For more information on how to use the library check out the docs:

[Event Bus Go Client - GoDoc](https://pkg.go.dev/github.com/go-web-dev/event-bus-go/client)

### Prerequisites

Before using the library make sure the Event Bus service is
up and running

```sh
# cd into the event bus project
cd event-bus

# compile and run the program
go run main.go
```

### Test

```sh
# run unit tests
go test ./...

# run tests with coverage
./coverage.sh

# run tests with coverage with generated HTML report
./coverage.sh --html
```
