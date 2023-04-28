package main

import "time"

// The option pattern is an excellent way to make your code more flexible and extensible.
// It allows users to create new objects with a custom configuration without having to add more parameters to the constructor.
// To summarize, the option pattern is a powerful tool for making code more flexible and extensible.
// It allows developers to add new configuration options to a constructor without cluttering it up with a ton of parameters.

type Server struct {
	host    string
	port    int
	timeout time.Duration
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.host = host
	}
}
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}
func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30 * time.Second,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
