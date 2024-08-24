package controllers

import (
	"context"
	"github.com/orewaee/bytebin/internal/app/api"
	"github.com/orewaee/bytebin/internal/handlers"
	"github.com/orewaee/bytebin/internal/middlewares"
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

	mux.HandleFunc("OPTIONS /*", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	mux.Handle("POST /bin", middlewares.CorsMiddleware(
		middlewares.LogMiddleware(handlers.NewPostHandler(bytebinApi, log), log)))

	mux.Handle("GET /bin/{id}", middlewares.CorsMiddleware(
		middlewares.LogMiddleware(handlers.NewGetHandler(bytebinApi, log), log)))

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
