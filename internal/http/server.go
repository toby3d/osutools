package http

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	S *http.Server
}

func NewServer(addr string, r *httprouter.Router) *Server {
	server := new(Server)
	server.S = &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         addr,
		Handler:      r,
	}

	return server
}
