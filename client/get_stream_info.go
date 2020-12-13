package client

import (
	"time"
)

type GetStreamInfoResponseBody struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
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

	var body GetStreamInfoResponseBody
	err = res.decodeBody(&body)
	if err != nil {
		return GetStreamInfoResponse{}, err
	}
	response := GetStreamInfoResponse{
		Response: res,
		Body: body,
	}

	return response, nil
}
