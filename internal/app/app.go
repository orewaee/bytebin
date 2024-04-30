package app

import (
	"github.com/orewaee/bytebin/internal/handlers"
	"log"
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
	Logger *log.Logger
}

func New(addr string, logger *log.Logger) *App {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthHandler)
	mux.HandleFunc("POST /bin", handlers.AddBin)
	mux.HandleFunc("GET /bin/{id}", handlers.GetBin)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return &App{
		Server: server,
		Logger: logger,
	}
}

func (app *App) Run() error {
	app.Logger.Println("running app at", app.Server.Addr)
	return app.Server.ListenAndServe()
}
