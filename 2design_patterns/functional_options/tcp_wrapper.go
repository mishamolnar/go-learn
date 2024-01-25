package functional_options

import (
	"net"
	"time"
)

//name: functional options
//design pattern that makes config optional and easier, without exposing private fields but with reaching them
//we could use setters, but then default config Would be nil value for each type
//also functional options makes possible to provide varargs in 1 line

type TcpServer struct {
	listener net.Listener
}

func WithRawMode() func(s *TcpServer) error {
	return func(s *TcpServer) error {
		//do something to set raw mode
		return nil
	}
}

func WithTimeout(t time.Duration) func(s *TcpServer) error {
	return func(s *TcpServer) error {
		//do something to specify timeout
		return nil
	}
}

func NewServer(addr string, opts ...func(s *TcpServer) error) (*TcpServer, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	s := TcpServer{l}
	for _, opt := range opts {
		err = opt(&s)
	}
	return &s, nil
}
