package client

import (
	"fmt"
)

func (s *clientSuite) Test_ProcessEvents_Success() {
	body := fmt.Sprintf(
		`{"events": [{"id": "%s", "stream_id": "%s", "status": 0, "created_at": "%s", "body": %s}]}`,
		"event-id",
		"stream-id",
		testTimeStr,
		`{"evtKey": "evtValue"}`,
	)
	expectedBody := ProcessEventsResponseBody{
		Events: []Event{
			{
				ID: "event-id",
				StreamID: "stream-id",
				Status: 0,
				CreatedAt: testTime,
				Body: []byte(`{"evtKey": "evtValue"}`),
			},
		},
	}
	expected := ProcessEventsResponse{
		Response: Response{
			Operation: "process_events",
			Status:    true,
		},
		Body: expectedBody,
	}
	s.writeRes("process_events", true, body, "")

	res, err := s.client.ProcessEvents("expenses")

	s.Require().NoError(err)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(expectedBody, res.Body)
	s.Nil(res.Reason)
}

func (s *clientSuite) Test_ProcessEvents_Failure() {
	reason := "could not process events"
	expected := ProcessEventsResponse{
		Response: Response{
			Operation: "process_events",
			Status:    false,
			Reason:    &reason,
		},
		Body: ProcessEventsResponseBody{},
	}
	s.writeRes("process_events", false, "", reason)

	res, err := s.client.ProcessEvents("expenses")

	s.EqualError(err, reason)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(&reason, res.Reason)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_ProcessEvents_JSONError() {
	s.write("}")

	res, err := s.client.ProcessEvents("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Response.Body)
}
