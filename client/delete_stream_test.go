package client

func (s *clientSuite) Test_DeleteStream_Success() {
	expected := Response{
		Operation: "delete_stream",
		Status: true,
	}
	s.writeRes("delete_stream", true, "", "")

	res, err := s.client.DeleteStream("expenses")

	s.Require().NoError(err)
	s.Equal(expected, res)
}

func (s *clientSuite) Test_DeleteStream_Failure() {
	reason := "could not delete stream"
	expected := Response{
		Operation: "delete_stream",
		Status: false,
		Reason: &reason,
	}
	s.writeRes("delete_stream", false, "", reason)

	res, err := s.client.DeleteStream("expenses")

	s.EqualError(err, reason)
	s.Equal(expected, res)
	s.Nil(res.Body)
}

func (s *clientSuite) Test_DeleteStream_JSONError() {
	s.write("}")

	res, err := s.client.DeleteStream("expenses")

	s.EqualError(err, "invalid character '}' looking for beginning of value")
	s.Empty(res)
	s.Nil(res.Body)
}
