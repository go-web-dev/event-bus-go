package client

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	testTimeStr = "2020-12-15T05:28:31.490416Z"
)

var (
	testTime, _ = time.Parse(time.RFC3339, testTimeStr)
)

type clientSuite struct {
	suite.Suite
	li     net.Listener
	conn   net.Conn
	client *Client
}

func (s *clientSuite) SetupSuite() {
	addr := "localhost:8080"
	credentials := Credentials{
		ClientID:     "some-client-id",
		ClientSecret: "some-client-secret",
	}

	li, err := net.Listen("tcp", addr)
	s.Require().NoError(err)
	s.li = li

	client, err := New(addr, credentials)
	s.Require().NoError(err)
	s.client = client

	conn, err := li.Accept()
	s.Require().NoError(err)
	s.conn = conn
}

func (s *clientSuite) TearDownSuite() {
	s.Require().NoError(s.li.Close())
}

func (s *clientSuite) Test_New() {
	credentials := Credentials{
		ClientID:     "some-client-id",
		ClientSecret: "some-client-secret",
	}

	client, err := New("localhost:8080", credentials)

	s.Require().NoError(err)
	s.Equal(credentials, client.credentials)
	s.NotNil(client.conn)
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

func TestClient(t *testing.T) {
	suite.Run(t, new(clientSuite))
}
