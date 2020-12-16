package client

func (s *clientSuite) Test_Exit_Success() {
	expected := Response{
		Operation: "exit",
		Status:    true,
	}
	s.writeRes("exit", true, "", "")

	res, err := s.client.Exit()

	s.Require().NoError(err)
	s.Equal(expected, res)
}

func (s *clientSuite) Test_Exit_Failure() {
	reason := "could not exit"
	expected := Response{
		Operation: "exit",
		Status:    false,
		Reason:    &reason,
	}
	s.writeRes("exit", false, "", reason)

	res, err := s.client.Exit()

	s.EqualError(err, reason)
	s.Equal(expected, res)
	s.Nil(res.Body)
}

func (s *clientSuite) Test_Exit_JSONReadError() {
	s.write("}")

	res, err := s.client.Exit()

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Body)
}
