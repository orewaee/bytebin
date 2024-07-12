package app

import (
	"github.com/orewaee/bytebin/internal/handlers"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type App struct {
	storage storage.Storage
	log     *zerolog.Logger
}

func New(storage storage.Storage, log *zerolog.Logger) *App {
	return &App{
		storage: storage,
		log:     log,
	}
}

func (app *App) Run(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("POST /bin", handlers.NewPostHandler(app.storage, app.log))
	mux.Handle("GET /bin/{id}", handlers.NewGetHandler(app.storage, app.log))

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	app.log.Info().Msgf("running app on addr %s", addr)
	return server.ListenAndServe()
}
