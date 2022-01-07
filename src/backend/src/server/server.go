package server

import (
	"net/http"

	"github.com/go-redis/redis/v8"
)

type Server struct {
	hostPort string
	redisDb *redis.Client // we have a redis client to call for db operations
}

func NewServer(hostPort string) *Server {
	return &Server{
		hostPort: hostPort,
		redisDb: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			DB:       0,  // use default DB
		}),
	}
}

func (s *Server) Run() error {
	muxServer := http.NewServeMux()
	muxServer.Handle("/items", http.HandlerFunc(s.crud))
	return http.ListenAndServe(s.hostPort, muxServer)
}

