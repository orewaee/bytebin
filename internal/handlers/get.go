package handlers

import (
	"github.com/orewaee/bytebin/internal/storage"
	"log"
	"net/http"
)

type GetHandler struct {
	storage storage.Storage
}

func NewGetHandler(storage storage.Storage) *GetHandler {
	return &GetHandler{
		storage: storage,
	}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	b, m, err := h.storage.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("not found")); err != nil {
			log.Println(err)
		}
		return
	}

	w.Header().Set("Content-Type", m.ContentType)
	if _, err := w.Write(b); err != nil {
		log.Println(err)
	}
}
