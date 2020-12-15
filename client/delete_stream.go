package client

import (
	"errors"
)

// DeleteStream deletes a stream along with all events associated to it
func (c Client) DeleteStream(name string) (Response, error) {
	r := req{
		Operation: "delete_stream",
		Body: map[string]string{
			"stream_name": name,
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
		return Response{}, errors.New(*res.Reason)
	}
	return res, nil
}
