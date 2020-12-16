package client

import (
	"errors"
)

type MarkEventRequestBody struct {
	EventID string `json:"event_id"`
	Status  int    `json:"status"`
}

// MarkEvent marks an event's status, changing its workflow i.e. after events processing or retrying
func (c Client) MarkEvent(eventID string, status int) (Response, error) {
	r := req{
		Operation: "mark_event",
		Body: MarkEventRequestBody{
			EventID: eventID,
			Status: status,
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
		return res, errors.New(*res.Reason)
	}

	return res, nil
}
