package client

import (
	"errors"
	"time"
)

// Stream represents short information about a stream
type Stream struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateStreamResponseBody represents the response body when creating a stream given by Event Bus
type CreateStreamResponseBody struct {
	Stream Stream `json:"stream"`
}

// CreateStreamResponse represents the response when creating a stream returned back from Event Bus
type CreateStreamResponse struct {
	Response
	Body CreateStreamResponseBody
}

// CreateStream creates a new empty stream available for specific events
func (c Client) CreateStream(name string) (CreateStreamResponse, error) {
	r := req{
		Operation: "create_stream",
		Body: map[string]string{
			"stream_name": name,
		},
	}

	err := c.write(r)
	if err != nil {
		return CreateStreamResponse{}, err
	}
	res, err := c.read()
	if err != nil {
		return CreateStreamResponse{}, err
	}
	if res.Reason != nil {
		r := CreateStreamResponse{}
		r.Operation = res.Operation
		r.Status = res.Status
		r.Reason = res.Reason
		return r, errors.New(*res.Reason)
	}

	var body CreateStreamResponseBody
	err = res.decodeBody(&body)
	if err != nil {
		return CreateStreamResponse{}, err
	}
	response := CreateStreamResponse{
		Response: res,
		Body:     body,
	}

	return response, nil
}
