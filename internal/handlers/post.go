package handlers

import (
	"github.com/orewaee/bytebin/internal/config"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/orewaee/bytebin/pkg/dto"
	"github.com/rs/xid"
	"io"
	"log"
	"net/http"
	"time"
)

type PostHandler struct {
	storage storage.Storage
}

func NewPostHandler(storage storage.Storage) *PostHandler {
	return &PostHandler{
		storage: storage,
	}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("failed to read bytes")); err != nil {
			log.Println(err)
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
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(id)); err != nil {
		log.Println(err)
	}
}
