package client

import (
	"fmt"
)

func (s *clientSuite) Test_GetStreamInfo_Success() {
	body := fmt.Sprintf(
		`{"stream": {"id": "stream-id", "name": "expenses", "created_at": "%s"}}`,
		testTimeStr,
	)
	expectedBody := GetStreamInfoResponseBody{
		Stream: Stream{
			ID:        "stream-id",
			Name:      "expenses",
			CreatedAt: testTime,
		},
	}
	expected := GetStreamInfoResponse{
		Response: Response{
			Operation: "get_stream_info",
			Status:    true,
		},
		Body: expectedBody,
	}
	s.writeRes("get_stream_info", true, body, "")

	res, err := s.client.GetStreamInfo("expenses")

	s.Require().NoError(err)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(expectedBody, res.Body)
	s.Nil(res.Reason)
}

func (s *clientSuite) Test_GetStreamInfo_Failure() {
	reason := "could not get stream info"
	expected := GetStreamInfoResponse{
		Response: Response{
			Operation: "get_stream_info",
			Status:    false,
			Reason:    &reason,
		},
		Body: GetStreamInfoResponseBody{},
	}
	s.writeRes("get_stream_info", false, "", reason)

	res, err := s.client.GetStreamInfo("expenses")

	s.EqualError(err, reason)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(&reason, res.Reason)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_GetStreamInfo_JSONError() {
	s.write("}")

	res, err := s.client.GetStreamInfo("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Response.Body)
}
