package client

func (s *clientSuite) Test_MarkEvent_Success() {
	expected := Response{
		Operation: "mark_event",
		Status: true,
	}
	s.writeRes("mark_event", true, "", "")

	res, err := s.client.MarkEvent("some-event-id", 1)

	s.Require().NoError(err)
	s.Equal(expected, res)
}

func (s *clientSuite) Test_MarkEvent_Failure() {
	reason := "could not mark event"
	expected := Response{
		Operation: "mark_event",
		Status: false,
		Reason: &reason,
	}
	s.writeRes("mark_event", false, "", reason)

	res, err := s.client.MarkEvent("some-event-id", 1)

	s.EqualError(err, reason)
	s.Equal(expected, res)
	s.Nil(res.Body)
}

func (s *clientSuite) Test_MarkEvent_JSONError() {
	s.write("}")

	res, err := s.client.MarkEvent("some-event-id", 1)

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Body)
}
