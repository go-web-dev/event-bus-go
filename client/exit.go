package client

func (c Client) Exit() (Response, error) {
	r := req{
		Operation: "exit",
	}
	err := c.write(r)
	if err != nil {
		return Response{}, nil
	}

	return c.read()
}
