package client

import (
	"errors"
)

// Exit tells the remote host Event Bus to close the current connection
func (c Client) Exit() (Response, error) {
	r := req{
		Operation: "exit",
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
		return res, errors.New(*res.Reason)
	}
	return res, nil
}
