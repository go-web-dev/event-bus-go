package client

import (
	"errors"
	"time"
)

type GetStreamInfoResponseBody struct {
	Stream struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"stream"`
}

type GetStreamInfoResponse struct {
	Response
	Body GetStreamInfoResponseBody
}

func (c Client) GetStreamInfo(name string) (GetStreamInfoResponse, error) {
	r := req{
		Operation: "get_stream_info",
		Body: map[string]string{
			"stream_name": name,
		},
	}

	err := c.write(r)
	if err != nil {
		return GetStreamInfoResponse{}, err
	}
	res, err := c.read()
	if err != nil {
		return GetStreamInfoResponse{}, err
	}
	if res.Reason != nil {
		return GetStreamInfoResponse{}, errors.New(*res.Reason)
	}

	var body GetStreamInfoResponseBody
	err = res.decodeBody(&body)
	if err != nil {
		return GetStreamInfoResponse{}, err
	}
	response := GetStreamInfoResponse{
		Response: res,
		Body:     body,
	}

	return response, nil
}
