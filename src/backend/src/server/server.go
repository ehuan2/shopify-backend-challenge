package server

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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

// split path - breaks up a url into the first part and the rest
func splitPath(p string) (string, string) {
    p = path.Clean("/" + p) // prepend a '/' and clean up anything
    i := strings.Index(p[1:], "/") + 1 // get the index of the next slash
    if i <= 0 { // if the index of 2nd slash doesn't exist, then we have the current path to be everything - beginning slash
			return p[1:], "/"
    }
    return p[1:i], p[i:] // otherwise, we can split the path up
}

func (s *Server) route(w http.ResponseWriter, r *http.Request) {
	// here, we'll route it to either /items or /items/{id}
	start, rest := splitPath(r.URL.Path)
	if start != "items" {
		http.Error(w, "route not found", http.StatusNotFound)
		return
	}

	// cors stuff
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// then we check if the rest exists (/) means that it doesn't
	if rest == "/" {
		s.allCrud(w, r) // let the others handle it
		return
	}

	// we set the id to be the rest without the beginning slash
	id := rest[1:]

	// then we try and parse rest as a uuid, error out otherwise
	_, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "route not found", http.StatusNotFound)
		return
	}

	// otherwise, let the individual route handle it
	s.crud(w, r, id)
}

func (s *Server) Run() error {
	muxServer := http.NewServeMux()
	muxServer.Handle("/", http.HandlerFunc(s.route))
	return http.ListenAndServe(s.hostPort, muxServer)
}

