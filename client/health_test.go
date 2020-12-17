package client

import (
	"time"
)

func (s *clientSuite) Test_Health_Success() {
	expected := Response{
		Operation: "health",
		Status:    true,
	}
	s.writeRes("health", true, "", "")

	res, err := s.client.Health()

	s.Require().NoError(err)
	s.Equal(expected, res)
}

func (s *clientSuite) Test_Health_Failure() {
	reason := "could not call for health"
	expected := Response{
		Operation: "health",
		Status:    false,
		Reason:    &reason,
	}
	s.writeRes("health", false, "", reason)

	res, err := s.client.Health()

	s.EqualError(err, reason)
	s.Equal(expected, res)
	s.Nil(res.Body)
}

func (s *clientSuite) Test_Health_JSONReadError() {
	s.write("}")

	res, err := s.client.Health()

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Body)
}

func (s *clientSuite) Test_Health_ConnReadError() {
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(-50 * time.Millisecond)))

	res, err := s.client.Health()

	s.Require().NotNil(err)
	s.Regexp("write tcp .* i/o timeout", err.Error())
	s.Empty(res)
}
