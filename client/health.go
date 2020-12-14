package client

import (
	"errors"
)

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
