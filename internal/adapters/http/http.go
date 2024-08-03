package http

import (
	"github.com/orewaee/bytebin/internal/app/ports"
	"log"
	"net/http"
	"time"
)

type Server struct {
	bytebinService ports.BytebinService
}

func NewServer(bytebinService ports.BytebinService) *Server {
	return &Server{
		bytebinService: bytebinService,
	}
}

func (server *Server) Run(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("POST /bin", NewPostHandler(server.bytebinService))
	mux.Handle("GET /bin/{id}", NewGetHandler(server.bytebinService))

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	log.Printf("running app on addr %s\n", addr)
	return srv.ListenAndServe()
}
