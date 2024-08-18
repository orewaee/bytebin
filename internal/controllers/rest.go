package controllers

import (
	"context"
	"github.com/orewaee/bytebin/internal/app/api"
	"github.com/orewaee/bytebin/internal/handlers"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type RestController struct {
	server     *http.Server
	bytebinApi api.BytebinApi
	log        *zerolog.Logger
}

func NewRestController(addr string, bytebinApi api.BytebinApi, log *zerolog.Logger) *RestController {
	mux := http.NewServeMux()

	mux.Handle("POST /bin", handlers.NewPostHandler(bytebinApi, log))
	mux.Handle("GET /bin/{id}", handlers.NewGetHandler(bytebinApi, log))

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	return &RestController{
		server:     server,
		bytebinApi: bytebinApi,
		log:        log,
	}
}

func (controller *RestController) Run() error {
	controller.log.Info().Msgf("running on addr %s", controller.server.Addr)
	return controller.server.ListenAndServe()
}

func (controller *RestController) Shutdown(ctx context.Context) error {
	return controller.server.Shutdown(ctx)
}