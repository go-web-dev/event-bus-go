package client

import (
	"encoding/json"
	"errors"
	"time"
)

type ProcessEventsResponseBody struct {
	Events []struct {
		ID        string          `json:"id"`
		StreamID  string          `json:"stream_id"`
		Name      string          `json:"name"`
		Body      json.RawMessage `json:"body"`
		CreatedAt time.Time       `json:"created_at"`
	} `json:"events"`
}

type ProcessEventsResponse struct {
	Response
	Body ProcessEventsResponseBody
}

func (c Client) ProcessEvents(name string) (ProcessEventsResponse, error) {
	r := req{
		Operation: "process_events",
		Body: map[string]string{
			"stream_name": name,
		},
	}

	err := c.write(r)
	if err != nil {
		return ProcessEventsResponse{}, err
	}
	res, err := c.read()
	if err != nil {
		return ProcessEventsResponse{}, err
	}
	if res.Reason != nil {
		return ProcessEventsResponse{}, errors.New(*res.Reason)
	}

	var body ProcessEventsResponseBody
	err = res.decodeBody(&body)
	if err != nil {
		return ProcessEventsResponse{}, err
	}
	response := ProcessEventsResponse{
		Response: res,
		Body:     body,
	}

	return response, nil
}
