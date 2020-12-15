package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"net"
)

type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Client struct {
	conn        *net.TCPConn
	credentials Credentials
}

func New(addr string, credentials Credentials) (*Client, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:        conn,
		credentials: credentials,
	}
	return c, nil
}

type req struct {
	Operation string      `json:"operation"`
	Body      interface{} `json:"body,omitempty"`
	Auth      Credentials `json:"auth"`
}

type Response struct {
	Operation string           `json:"operation"`
	Body      *json.RawMessage `json:"body,omitempty"`
	Status    bool             `json:"status"`
	Reason    *string          `json:"reason,omitempty"`
}

func (r Response) decodeBody(v interface{}) error {
	if r.Body == nil {
		return errors.New("cannot decode nil body")
	}
	return json.Unmarshal(*r.Body, v)
}

func (c Client) write(r req) error {
	r.Auth = c.credentials
	bs, err := json.Marshal(r)
	if err != nil {
		return err
	}

	bs = append(bs, '\n')
	_, err = c.conn.Write(bs)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) read() (Response, error) {
	reader := bufio.NewReader(c.conn)
	bs, err := reader.ReadBytes('\n')
	if err != nil {
		return Response{}, err
	}

	var res Response
	err = json.Unmarshal(bs, &res)
	if err != nil {
		return Response{}, err
	}

	return res, nil
}
