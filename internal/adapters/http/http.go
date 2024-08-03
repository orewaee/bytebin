package http

import (
	"github.com/orewaee/bytebin/internal/app/ports"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type Server struct {
	bytebinService ports.BytebinService
	log            *zerolog.Logger
}

func NewServer(bytebinService ports.BytebinService, log *zerolog.Logger) *Server {
	return &Server{
		bytebinService: bytebinService,
		log:            log,
	}
}

func (server *Server) Run(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("POST /bin", NewPostHandler(server.bytebinService, server.log))
	mux.Handle("GET /bin/{id}", NewGetHandler(server.bytebinService, server.log))

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	server.log.Info().Msgf("running app on addr %s", addr)
	return srv.ListenAndServe()
}
