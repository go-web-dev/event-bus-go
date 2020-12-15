package client

import (
	"encoding/json"
	"errors"
	"time"
)

// ProcessEventsResponseBody represents the response body when processing available events on a stream inside Event Bus
type ProcessEventsResponseBody struct {
	Events []struct {
		ID        string          `json:"id"`
		StreamID  string          `json:"stream_id"`
		Status    int             `json:"status"`
		Name      string          `json:"name"`
		Body      json.RawMessage `json:"body"`
		CreatedAt time.Time       `json:"created_at"`
	} `json:"events"`
}

// ProcessEventsResponseBody represents the response when processing available events on a stream inside Event Bus
type ProcessEventsResponse struct {
	Response
	Body ProcessEventsResponseBody
}

// ProcessEvents fetches all events marked as unprocessed from a specific stream
func (c Client) ProcessEvents(name string) (ProcessEventsResponse, error) {
	return c.events(name, "process_events")
}

// ProcessEvents fetches all events marked as retry from a specific stream
func (c Client) RetryEvents(name string) (ProcessEventsResponse, error) {
	return c.events(name, "retry_events")
}

// ProcessEvents fetches all events from a specific stream regardless of their statuses
func (c Client) GetStreamEvents(name string) (ProcessEventsResponse, error) {
	return c.events(name, "get_stream_events")
}

func (c Client) events(streamName, operation string) (ProcessEventsResponse, error) {
	r := req{
		Operation: operation,
		Body: map[string]string{
			"stream_name": streamName,
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
