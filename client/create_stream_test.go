package client

import (
	"fmt"
	"time"
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
	s.writeRes("create_stream", true, body, "")

	res, err := s.client.CreateStream("expenses")

	s.Require().NoError(err)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(expectedBody, res.Body)
	s.Nil(res.Reason)
}

func (s *clientSuite) Test_CreateStream_Failure() {
	reason := "could not create stream"
	expected := CreateStreamResponse{
		Response: Response{
			Operation: "create_stream",
			Status:    false,
			Reason:    &reason,
		},
		Body: CreateStreamResponseBody{},
	}
	s.writeRes("create_stream", false, "", reason)

	res, err := s.client.CreateStream("expenses")

	s.EqualError(err, reason)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(&reason, res.Reason)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_CreateStream_JSONReadError() {
	s.write("}")

	res, err := s.client.CreateStream("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_CreateStream_ConnReadError() {
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(-50 * time.Millisecond)))

	res, err := s.client.CreateStream("expenses")

	s.Require().NotNil(err)
	s.Regexp("write tcp .* i/o timeout", err.Error())
	s.Empty(res)
}

func (s *clientSuite) Test_CreateStream_JSONDecodeError() {
	s.write(`{}`)

	res, err := s.client.CreateStream("expenses")

	s.EqualError(err, "cannot decode nil body")
	s.Empty(res)
	s.Nil(res.Response.Body)
}
