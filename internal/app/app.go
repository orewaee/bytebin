package app

import (
	"github.com/orewaee/bytebin/internal/handlers"
	"github.com/orewaee/bytebin/internal/storage"
	"log"
	"net/http"
	"time"
)

type App struct {
	storage storage.Storage
}

func New(storage storage.Storage) *App {
	return &App{
		storage: storage,
	}
}

func (app *App) Run(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("POST /bin", handlers.NewPostHandler(app.storage))
	mux.Handle("GET /bin/{id}", handlers.NewGetHandler(app.storage))

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	log.Println("running app on addr " + addr)
	return server.ListenAndServe()
}
