package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

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

type Client struct {
	conn *net.TCPConn
}

type Response struct {
	Operation string           `json:"operation"`
	Body      *json.RawMessage `json:"body,omitempty"`
	Status    bool             `json:"status"`
	Reason    *string           `json:"reason,omitempty"`
}

type req struct {
	Operation string      `json:"operation"`
	Body      interface{} `json:"body,omitempty"`
}

func (c Client) HealthOperation() (Response, error) {
	op := req{
		Operation: "health",
	}
	bs, err := json.Marshal(op)
	if err != nil {
		fmt.Println("could not marshal operation")
		return Response{}, err
	}
	bs = append(bs, '\n')
	_, err = c.conn.Write(bs)
	if err != nil {
		fmt.Println("could not write to connection")
		return Response{}, err
	}

	reader := bufio.NewReader(c.conn)
	bs, err = reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("could not read response from connection")
		return Response{}, err
	}

	var res Response
	err = json.Unmarshal(bs, &res)
	if err != nil {
		fmt.Println("could not unmarshal response")
		return Response{}, err
	}

	return res, nil
}
