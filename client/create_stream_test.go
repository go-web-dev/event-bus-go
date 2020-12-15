package client

import (
	"fmt"
)

func (s *clientSuite) Test_CreateStream_Success() {
	body := fmt.Sprintf(
		`{"stream": {"id": "stream-id", "name": "expenses", "created_at": "%s"}}`,
		testTimeStr,
	)
	expectedBody := CreateStreamResponseBody{
		Stream: Stream{
			ID:        "stream-id",
			Name:      "expenses",
			CreatedAt: testTime,
		},
	}
	expected := CreateStreamResponse{
		Response: Response{
			Operation: "create_stream",
			Status:    true,
		},
		Body: expectedBody,
	}

	s.write("create_stream", true, body, "")
	res, err := s.client.CreateStream("expenses")

	s.Require().NoError(err)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(expectedBody, res.Body)
	s.Nil(res.Reason)
}

func (s *clientSuite) Test_CreateStream_Error() {
	reason := `could not create stream`
	expected := CreateStreamResponse{
		Response: Response{
			Operation: "create_stream",
			Status:    false,
			Reason: &reason,
		},
		Body: CreateStreamResponseBody{},
	}

	s.write("create_stream", false, "", reason)
	res, err := s.client.CreateStream("expenses")

	s.EqualError(err, "could not create stream")
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(&reason, res.Reason)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_CreateStream_JSONError() {
	_, err := s.conn.Write([]byte("}\n"))
	s.Require().NoError(err)

	res, err := s.client.CreateStream("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Response.Body)
}
