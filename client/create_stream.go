package client

import (
	"errors"
	"time"
)

type CreateStreamResponseBody struct {
	Stream struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"stream"`
}

type CreateStreamResponse struct {
	Response
	Body CreateStreamResponseBody
}

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
		return CreateStreamResponse{}, errors.New(*res.Reason)
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
