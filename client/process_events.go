package client

import (
	"fmt"
	"time"
)

type ProcessEventsResponseBody struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ProcessEventsResponse struct {
	Response
	Body []ProcessEventsResponseBody
}

func (c Client) ProcessEvents(name string) (ProcessEventsResponse, error) {
	r := req{
		Operation: "process_events",
		Body: map[string]string{
			"stream_name": name,
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
	fmt.Println(string(*res.Body))

	var body []ProcessEventsResponseBody
	err = res.decodeBody(&body)
	if err != nil {
		return ProcessEventsResponse{}, err
	}
	response := ProcessEventsResponse{
		Response: res,
		Body: body,
	}

	return response, nil
}
