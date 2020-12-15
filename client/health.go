package client

import (
	"errors"
)

// Health responds back with a successful response to indicate the Event Bus service is uo and running
func (c Client) Health() (Response, error) {
	r := req{
		Operation: "health",
	}
	err := c.write(r)
	if err != nil {
		return Response{}, nil
	}

	res, err := c.read()
	if err != nil {
		return Response{}, err
	}
	if res.Reason != nil {
		return Response{}, errors.New(*res.Reason)
	}
	return res, nil
}
