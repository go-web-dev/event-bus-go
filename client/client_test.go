package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	testTimeStr = "2020-12-15T05:28:31.490416Z"
)

var (
	addr        = "localhost:8080"
	testTime, _ = time.Parse(time.RFC3339, testTimeStr)
	credentials = Credentials{
		ClientID:     "some-client-id",
		ClientSecret: "some-client-secret",
	}
)

type clientSuite struct {
	suite.Suite
	li     net.Listener
	conn   net.Conn
	client *Client
	wg     sync.WaitGroup
}

func (s *clientSuite) SetupTest() {
	li, err := net.Listen("tcp", addr)
	s.Require().NoError(err)
	s.li = li

	client, err := New(addr, credentials)
	s.Require().NoError(err)
	s.client = client

	s.wg.Add(1)
	go func() {
		conn, err := s.li.Accept()
		s.Require().NoError(err)
		s.conn = conn
		s.wg.Done()
	}()
	s.wg.Wait()
}

func (s *clientSuite) TearDownTest() {
	s.Require().NoError(s.li.Close())
	s.Require().NoError(s.conn.Close())
}

func (s *clientSuite) Test_New_Success() {
	client, err := New(addr, credentials)

	s.Require().NoError(err)
	s.Equal(credentials, client.credentials)
	s.NotNil(client.conn)
}

func (s *clientSuite) Test_New_TCPAddrError() {
	client, err := New("hello", credentials)

	s.EqualError(err, "address hello: missing port in address")
	s.Nil(client)
}

func (s *clientSuite) Test_New_DialTCPError() {
	client, err := New("localhost:9000", credentials)

	s.EqualError(err, "dial tcp 127.0.0.1:9000: connect: connection refused")
	s.Nil(client)
}

func (s *clientSuite) Test_write_Success() {
	expected := `{"operation":"some-operation","body":{"k":"v"},"auth":{"client_id":"some-client-id","client_secret":"some-client-secret"}}`
	expected += "\n"
	r := req{
		Operation: "some-operation",
		Auth:      credentials,
		Body: map[string]string{
			"k": "v",
		},
	}

	err := s.client.write(r)

	s.Require().NoError(err)
}

func (s *clientSuite) Test_write_Error() {
	r := req{
		Operation: "some-operation",
		Auth:      credentials,
		Body: map[string]string{
			"k": "v",
		},
	}
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(-50 * time.Millisecond)))

	err := s.client.write(r)

	s.Require().NotNil(err)
	s.Regexp("write tcp .* i/o timeout", err.Error())
}

func (s *clientSuite) Test_read_Success() {
	resJSON := `{"operation":"some-operation", "status": true, "body":{"k":"v"}}`
	resJSON += "\n"
	body := json.RawMessage(`{"k":"v"}`)
	expected := Response{
		Operation: "some-operation",
		Status:    true,
		Body:      &body,
	}
	s.write(resJSON)

	res, err := s.client.read()

	s.Require().NoError(err)
	s.Equal(expected, res)
}

func (s *clientSuite) Test_read_ConnReadError() {
	s.Require().NoError(s.client.conn.SetDeadline(time.Now().Add(50 * time.Millisecond)))

	res, err := s.client.read()

	s.Require().NotNil(err)
	s.Regexp("read tcp .* i/o timeout", err.Error())
	s.Empty(res)
}

func (s *clientSuite) Test_read_UnmarshalError() {
	resJSON := `{"operation": 1}` + "\n"
	s.write(resJSON)

	res, err := s.client.read()

	s.EqualError(err, "json: cannot unmarshal number into Go struct field Response.operation of type string")
	s.Empty(res)
}

func (s *clientSuite) read(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	bs, err := reader.ReadBytes('\n')
	s.Require().NoError(err)
	return string(bs)
}

func (s *clientSuite) write(text string) {
	_, err := s.conn.Write([]byte(text + "\n"))
	s.Require().NoError(err)
}

func (s *clientSuite) writeRes(op string, status bool, body, reason string) {
	res := fmt.Sprintf(`{"operation": "%s", "status": %t`, op, status)
	if body != "" {
		res += fmt.Sprintf(`, "body": %s`, body)
	}
	if reason != "" {
		res += fmt.Sprintf(`, "reason": "%s"`, reason)
	}
	res += "}"
	s.write(res)
}

func (s *clientSuite) checkConn() {
	// check that nobody writes to conn
}

func TestClient(t *testing.T) {
	suite.Run(t, new(clientSuite))
}
