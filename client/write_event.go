package client

import (
	"errors"
)

// WriteEventRequestBody represents request body for creating an event in the Event Bus
type WriteEventRequestBody struct {
	StreamName string      `json:"stream_name"`
	Event      interface{} `json:"event"`
}

// WriteEvent writes an event to an existing stream
func (c Client) WriteEvent(streamName string, event interface{}) (Response, error) {
	r := req{
		Operation: "write_event",
		Body: WriteEventRequestBody{
			StreamName: streamName,
			Event:      event,
		},
	}

	err := c.write(r)
	if err != nil {
		return Response{}, err
	}
	res, err := c.read()
	if err != nil {
		return Response{}, err
	}
	if res.Reason != nil {
		return res, errors.New(*res.Reason)
	}

	return res, nil
}
