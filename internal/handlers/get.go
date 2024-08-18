package handlers

import (
	"github.com/orewaee/bytebin/internal/app/api"
	"github.com/rs/zerolog"
	"net/http"
)

type GetHandler struct {
	bytebinApi api.BytebinApi
	log        *zerolog.Logger
}

func NewGetHandler(bytebinApi api.BytebinApi, log *zerolog.Logger) *GetHandler {
	return &GetHandler{
		bytebinApi: bytebinApi,
		log:        log,
	}
}

func (handler *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.log.Debug().
		Str("method", "GET").
		Str("route", r.URL.String()).
		Send()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("id")

	bin, meta, err := handler.bytebinApi.GetById(id)
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
