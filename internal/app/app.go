package app

import (
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/rs/xid"
	"io"
	"log"
	"net/http"
	"time"
)

type App struct {
	Addr   string
	Server *http.Server
}

func New(addr string) *App {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /bin", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		contentType := r.Header.Get("Content-Type")

		id := xid.New().String()

		storage.AddBin(id, &bin.Bin{
			ContentType: contentType,
			Bytes:       bytes,
		})

		w.WriteHeader(http.StatusCreated)

		_, err = w.Write([]byte(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	mux.HandleFunc("GET /bin/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		b, ok := storage.GetBin(id)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Header().Set("Content-Type", b.ContentType)
		_, err := w.Write(b.Bytes)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return &App{
		Addr:   addr,
		Server: server,
	}
}

func (app *App) Run() error {
	log.Println("running app on", app.Addr)
	if err := app.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
