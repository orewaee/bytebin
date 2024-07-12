package handlers

import (
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/rs/zerolog"
	"net/http"
)

type GetHandler struct {
	storage storage.Storage
	log     *zerolog.Logger
}

func NewGetHandler(storage storage.Storage, log *zerolog.Logger) *GetHandler {
	return &GetHandler{
		storage: storage,
		log:     log,
	}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Debug().Str("method", "GET").Str("url", r.URL.String()).Send()

	id := r.PathValue("id")

	b, m, err := h.storage.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("not found")); err != nil {
			h.log.Error().Err(err).Send()
		}
		return
	}

	w.Header().Set("Content-Type", m.ContentType)
	if _, err := w.Write(b); err != nil {
		h.log.Error().Err(err).Send()
	}
}
