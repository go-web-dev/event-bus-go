package client

import "time"

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

func (s *clientSuite) Test_WriteEvent_JSONReadError() {
	s.write("}")

	res, err := s.client.WriteEvent("expenses", testEvt)

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Body)
}

func (s *clientSuite) Test_WriteEvent_ConnReadError() {
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(-50 * time.Millisecond)))

	res, err := s.client.WriteEvent("expenses", nil)

	s.Require().NotNil(err)
	s.Regexp("write tcp .* i/o timeout", err.Error())
	s.Empty(res)
}
