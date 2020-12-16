package client

import (
	"errors"
)

// GetStreamInfoResponseBody represents the response body when getting the info about a stream from Event Bus
type GetStreamInfoResponseBody struct {
	Stream Stream `json:"stream"`
}

// GetStreamInfoResponseBody represents the response when getting the info about a stream from Event Bus
type GetStreamInfoResponse struct {
	Response
	Body GetStreamInfoResponseBody
}

// GetStreamInfo gets the short info about a created stream
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
		r := GetStreamInfoResponse{}
		r.Operation = res.Operation
		r.Status = res.Status
		r.Reason = res.Reason
		return r, errors.New(*res.Reason)
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
