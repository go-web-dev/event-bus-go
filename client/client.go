package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"net"
)

// Credentials represents the credentials for the client using the library.
// They have to match the credentials stored inside Event Bus for that specific client
type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Client represents the Event Bus client that communicates over TCP according to a specific protocol.
// The connection that is created is only closed when calling Exit() or if the remote host closes it
type Client struct {
	conn        *net.TCPConn
	credentials Credentials
}

// New creates a brand new Event Bus client that communicates TCP with the Event Bus service
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

// Response represents the generic response for every request returned by Event Bus service.
// Operation and Status are always present.
// Body is present only in case of a success operation, otherwise it's nil and Status is false
// Reason is present only in case of a failed operation, otherwise it's nil and Status is true
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
