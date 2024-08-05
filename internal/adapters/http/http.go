package http

import (
	"context"
	"github.com/orewaee/bytebin/internal/app/ports"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type Server struct {
	srv            *http.Server
	bytebinService ports.BytebinService
	log            *zerolog.Logger
}

func NewServer(bytebinService ports.BytebinService, log *zerolog.Logger) *Server {
	mux := http.NewServeMux()

	mux.Handle("POST /bin", NewPostHandler(bytebinService, log))
	mux.Handle("GET /bin/{id}", NewGetHandler(bytebinService, log))

	srv := &http.Server{
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	return &Server{
		srv:            srv,
		bytebinService: bytebinService,
		log:            log,
	}
}

func (server *Server) Run(addr string) error {
	server.srv.Addr = addr
	server.log.Info().Msgf("running server on addr %s...", addr)
	return server.srv.ListenAndServe()
}

func (server *Server) Shutdown() error {
	server.log.Info().Msg("stopping the server...")
	return server.srv.Shutdown(context.TODO())
}
