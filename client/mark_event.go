package client

import (
	"errors"
)

func (c Client) MarkEvent(eventID string, status int) (Response, error) {
	r := req{
		Operation: "mark_event",
		Body: map[string]interface{}{
			"event_id": eventID,
			"status":   status,
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
