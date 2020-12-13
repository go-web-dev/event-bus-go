package client

func (c Client) DeleteStream(name string) (Response, error) {
	r := req{
		Operation: "delete_stream",
		Body: map[string]string{
			"stream_name": name,
		},
	}
	err := c.write(r)
	if err != nil {
		return Response{}, err
	}

	return c.read()
}
