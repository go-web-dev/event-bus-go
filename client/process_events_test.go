package client

import (
	"fmt"
	"time"
)

func (s *clientSuite) Test_ProcessEvents_Success() {
	body := fmt.Sprintf(
		`{"events": [%s]}`,
		s.eventJSON("event-id", "stream-id", `{"evtKey": "evtValue"}`, testTimeStr, 0),
	)
	expectedBody := ProcessEventsResponseBody{
		Events: []Event{
			{
				ID:        "event-id",
				StreamID:  "stream-id",
				Status:    0,
				CreatedAt: testTime,
				Body:      []byte(`{"evtKey": "evtValue"}`),
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

func (s *clientSuite) Test_ProcessEvents_JSONReadError() {
	s.write("}")

	res, err := s.client.ProcessEvents("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_ProcessEvents_JSONDecodeError() {
	s.write(`{}`)

	res, err := s.client.ProcessEvents("expenses")

	s.EqualError(err, "cannot decode nil body")
	s.Empty(res)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_ProcessEvents_ConnReadError() {
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(-50 * time.Millisecond)))

	res, err := s.client.ProcessEvents("expenses")

	s.Require().NotNil(err)
	s.Regexp("write tcp .* i/o timeout", err.Error())
	s.Empty(res)
}

func (s *clientSuite) Test_RetryEvents_Success() {
	body := fmt.Sprintf(
		`{"events": [%s]}`,
		s.eventJSON("event-id", "stream-id", `{"evtKey": "evtValue"}`, testTimeStr, 2),
	)
	expectedBody := ProcessEventsResponseBody{
		Events: []Event{
			{
				ID:        "event-id",
				StreamID:  "stream-id",
				Status:    2,
				CreatedAt: testTime,
				Body:      []byte(`{"evtKey": "evtValue"}`),
			},
		},
	}
	expected := ProcessEventsResponse{
		Response: Response{
			Operation: "retry_events",
			Status:    true,
		},
		Body: expectedBody,
	}
	s.writeRes("retry_events", true, body, "")

	res, err := s.client.RetryEvents("expenses")

	s.Require().NoError(err)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(expectedBody, res.Body)
	s.Nil(res.Reason)
}

func (s *clientSuite) Test_RetryEvents_Failure() {
	reason := "could not retry events"
	expected := ProcessEventsResponse{
		Response: Response{
			Operation: "retry_events",
			Status:    false,
			Reason:    &reason,
		},
		Body: ProcessEventsResponseBody{},
	}
	s.writeRes("retry_events", false, "", reason)

	res, err := s.client.RetryEvents("expenses")

	s.EqualError(err, reason)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(&reason, res.Reason)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_RetryEvents_JSONReadError() {
	s.write("}")

	res, err := s.client.RetryEvents("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_RetryEvents_JSONDecodeError() {
	s.write(`{}`)

	res, err := s.client.RetryEvents("expenses")

	s.EqualError(err, "cannot decode nil body")
	s.Empty(res)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_RetryEvents_ConnReadError() {
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(-50 * time.Millisecond)))

	res, err := s.client.RetryEvents("expenses")

	s.Require().NotNil(err)
	s.Regexp("write tcp .* i/o timeout", err.Error())
	s.Empty(res)
}

func (s *clientSuite) Test_GetStreamEvents_Success() {
	body := fmt.Sprintf(
		`{"events": [%s, %s, %s]}`,
		s.eventJSON("event-id1", "stream-id", `{"evtKey1": "evtValue1"}`, testTimeStr, 0),
		s.eventJSON("event-id2", "stream-id", `{"evtKey2": "evtValue2"}`, testTimeStr, 1),
		s.eventJSON("event-id3", "stream-id", `{"evtKey3": "evtValue3"}`, testTimeStr, 2),
	)
	expectedBody := ProcessEventsResponseBody{
		Events: []Event{
			{
				ID:        "event-id1",
				StreamID:  "stream-id",
				Status:    0,
				CreatedAt: testTime,
				Body:      []byte(`{"evtKey1": "evtValue1"}`),
			},
			{
				ID:        "event-id2",
				StreamID:  "stream-id",
				Status:    1,
				CreatedAt: testTime,
				Body:      []byte(`{"evtKey2": "evtValue2"}`),
			},
			{
				ID:        "event-id3",
				StreamID:  "stream-id",
				Status:    2,
				CreatedAt: testTime,
				Body:      []byte(`{"evtKey3": "evtValue3"}`),
			},
		},
	}
	expected := ProcessEventsResponse{
		Response: Response{
			Operation: "get_stream_events",
			Status:    true,
		},
		Body: expectedBody,
	}
	s.writeRes("get_stream_events", true, body, "")

	res, err := s.client.GetStreamEvents("expenses")

	s.Require().NoError(err)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(expectedBody, res.Body)
	s.Nil(res.Reason)
}

func (s *clientSuite) Test_GetStreamEvents_Failure() {
	reason := "could not get stream events"
	expected := ProcessEventsResponse{
		Response: Response{
			Operation: "get_stream_events",
			Status:    false,
			Reason:    &reason,
		},
		Body: ProcessEventsResponseBody{},
	}
	s.writeRes("get_stream_events", false, "", reason)

	res, err := s.client.GetStreamEvents("expenses")

	s.EqualError(err, reason)
	s.Equal(expected.Operation, res.Operation)
	s.Equal(expected.Status, res.Status)
	s.Equal(&reason, res.Reason)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_GetStreamEvents_JSONReadError() {
	s.write("}")

	res, err := s.client.GetStreamEvents("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_GetStreamEvents_JSONDecodeError() {
	s.write(`{}`)

	res, err := s.client.GetStreamEvents("expenses")

	s.EqualError(err, "cannot decode nil body")
	s.Empty(res)
	s.Nil(res.Response.Body)
}

func (s *clientSuite) Test_GetStreamEvents_ConnReadError() {
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(-50 * time.Millisecond)))

	res, err := s.client.GetStreamEvents("expenses")

	s.Require().NotNil(err)
	s.Regexp("write tcp .* i/o timeout", err.Error())
	s.Empty(res)
}

func (s *clientSuite) eventJSON(eventID, streamID, body, createdAt string, status uint8) string {
	return fmt.Sprintf(
		`{"id": "%s", "stream_id": "%s", "status": %d, "created_at": "%s", "body": %s}`,
		eventID,
		streamID,
		status,
		createdAt,
		body,
	)
}
