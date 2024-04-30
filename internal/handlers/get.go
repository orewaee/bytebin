package handlers

import (
	"github.com/orewaee/bytebin/internal/storage"
	"net/http"
)

func GetBin(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	b, ok := storage.GetBin(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", b.ContentType)
	if _, err := w.Write(b.Bytes); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
