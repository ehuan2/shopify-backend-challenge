package csv

import (
	"net/http"
)

// this folder provides the microservice to create the csv files and such

type Server struct {
	hostPort string
}

func NewServer(hostPort string) *Server {
	return &Server{
		hostPort: hostPort,
	}
}


func (s *Server) Run() error {
	muxServer := http.NewServeMux()
	muxServer.Handle("/", http.HandlerFunc(s.csv))
	return http.ListenAndServe(s.hostPort, muxServer)
}


