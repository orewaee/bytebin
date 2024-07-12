package handlers

import (
	"github.com/orewaee/bytebin/internal/config"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/orewaee/bytebin/pkg/dto"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"time"
)

type PostHandler struct {
	storage storage.Storage
	log     *zerolog.Logger
}

func NewPostHandler(storage storage.Storage, log *zerolog.Logger) *PostHandler {
	return &PostHandler{
		storage: storage,
		log:     log,
	}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Debug().Str("method", "POST").Str("url", r.URL.String()).Send()

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("failed to read bytes")); err != nil {
			h.log.Error().Err(err).Msg("failed to write response")
		}
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(bytes)
	}

	id := xid.New().String()

	ip := r.RemoteAddr
	userAgent := r.UserAgent()

	m := &dto.Meta{
		Id:          id,
		ContentType: contentType,
		Ip:          ip,
		UserAgent:   userAgent,
		CreatedAt:   time.Now(),
		Lifetime:    config.Get().Lifetime,
	}

	if err := h.storage.Add(id, bytes, m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			h.log.Error().Err(err).Msg("failed to write response")
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(id)); err != nil {
		h.log.Error().Err(err).Msg("failed to write response")
	}
}
