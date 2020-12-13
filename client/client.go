package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"net"
)

type Client struct {
	conn *net.TCPConn
}

func New(addr string) (*Client, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn: conn,
	}
	return c, nil
}

type req struct {
	Operation string      `json:"operation"`
	Body      interface{} `json:"body,omitempty"`
}

type Response struct {
	Operation string  `json:"operation"`
	Body      *json.RawMessage   `json:"body,omitempty"`
	Status    bool    `json:"status"`
	Reason    *string `json:"reason,omitempty"`
}

func (r Response) decodeBody(v interface{}) error {
	if r.Body == nil {
		return errors.New("cannot decode nil body")
	}
	return json.Unmarshal(*r.Body, v)
}

func (c Client) write(r req) error {
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
