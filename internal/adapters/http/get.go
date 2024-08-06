package http

import (
	"github.com/orewaee/bytebin/internal/app/ports"
	"github.com/rs/zerolog"
	"net/http"
)

type GetHandler struct {
	bytebinService ports.BytebinService
	log            *zerolog.Logger
}

func NewGetHandler(bytebinService ports.BytebinService, log *zerolog.Logger) *GetHandler {
	return &GetHandler{
		bytebinService: bytebinService,
		log:            log,
	}
}

func (handler *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.log.Debug().
		Str("method", "GET").
		Str("route", r.URL.String()).
		Send()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("id")

	bin, meta, err := handler.bytebinService.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		message := []byte("not found")
		if _, err := w.Write(message); err != nil {
			handler.log.Err(err).Send()
		}

		return
	}

	w.Header().Set("Content-Type", meta.ContentType)
	if _, err := w.Write(bin); err != nil {
		handler.log.Err(err).Send()
	}
}
