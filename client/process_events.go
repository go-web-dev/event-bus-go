package client

import (
	"encoding/json"
	"errors"
	"time"
)

// Event represents the event inside the Event Bus
type Event struct {
	ID        string          `json:"id"`
	StreamID  string          `json:"stream_id"`
	Status    uint8           `json:"status"`
	Body      json.RawMessage `json:"body"`
	CreatedAt time.Time       `json:"created_at"`
}

// ProcessEventsResponseBody represents the response body when processing available events on a stream inside Event Bus
type ProcessEventsResponseBody struct {
	Events []Event `json:"events"`
}

// ProcessEventsResponse represents the response when processing available events on a stream inside Event Bus
type ProcessEventsResponse struct {
	Response
	Body ProcessEventsResponseBody
}

// ProcessEvents fetches all events marked as unprocessed from a specific stream
func (c Client) ProcessEvents(streamName string) (ProcessEventsResponse, error) {
	return c.events(streamName, "process_events")
}

// RetryEvents fetches all events marked as retry from a specific stream
func (c Client) RetryEvents(streamName string) (ProcessEventsResponse, error) {
	return c.events(streamName, "retry_events")
}

// GetStreamEvents fetches all events from a specific stream regardless of their statuses
func (c Client) GetStreamEvents(streamName string) (ProcessEventsResponse, error) {
	return c.events(streamName, "get_stream_events")
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
		r := ProcessEventsResponse{}
		r.Operation = res.Operation
		r.Status = res.Status
		r.Reason = res.Reason
		return r, errors.New(*res.Reason)
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
