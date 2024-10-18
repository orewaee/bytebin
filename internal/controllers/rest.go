package controllers

import (
	"context"
	"github.com/orewaee/bytebin/internal/app/api"
	"github.com/orewaee/bytebin/internal/middlewares"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type RestController struct {
	server     *http.Server
	bytebinApi api.BytebinApi
	logger     *zerolog.Logger
}

func NewRestController(addr string, bytebinApi api.BytebinApi, logger *zerolog.Logger) *RestController {
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	return &RestController{
		server:     server,
		bytebinApi: bytebinApi,
		logger:     logger,
	}
}

func (controller *RestController) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("OPTIONS /*", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	mux.Handle("POST /bin", middlewares.CorsMiddleware(
		middlewares.LogMiddleware(controller.post, controller.logger)))

	mux.Handle("GET /bin/{id}", middlewares.CorsMiddleware(
		middlewares.LogMiddleware(controller.get, controller.logger)))

	controller.server.Handler = mux

	controller.logger.Info().Msgf("running on addr %s", controller.server.Addr)
	return controller.server.ListenAndServe()
}

func (controller *RestController) Shutdown(ctx context.Context) error {
	controller.logger.Info().Msg("shutting down...")
	return controller.server.Shutdown(ctx)
}
