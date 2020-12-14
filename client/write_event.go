package client

import "errors"

func (c Client) WriteEvent(name string, event interface{}) (Response, error) {
	r := req{
		Operation: "write_event",
		Body: map[string]interface{}{
			"stream_name": name,
			"event":       event,
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
