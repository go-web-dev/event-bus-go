package client

func (c Client) Health() (Response, error) {
	r := req{
		Operation: "health",
	}
	err := c.write(r)
	if err != nil {
		return Response{}, nil
	}

	return c.read()
}
