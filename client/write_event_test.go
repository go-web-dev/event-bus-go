package client

var (
	testEvt = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
)

func (s *clientSuite) Test_WriteEvent_Success() {
	expected := Response{
		Operation: "write_event",
		Status:    true,
	}
	s.writeRes("write_event", true, "", "")

	res, err := s.client.WriteEvent("expenses", testEvt)

	s.Require().NoError(err)
	s.Equal(expected, res)
}

func (s *clientSuite) Test_WriteEvent_Failure() {
	reason := "could not write event"
	expected := Response{
		Operation: "write_event",
		Status:    false,
		Reason:    &reason,
	}
	s.writeRes("write_event", false, "", reason)

	res, err := s.client.WriteEvent("expenses", testEvt)

	s.EqualError(err, reason)
	s.Equal(expected, res)
	s.Nil(res.Body)
}

func (s *clientSuite) Test_WriteEvent_JSONError() {
	s.write("}")

	res, err := s.client.WriteEvent("expenses", testEvt)

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Body)
}
