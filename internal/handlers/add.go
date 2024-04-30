package handlers

import (
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/rs/xid"
	"io"
	"net/http"
)

func AddBin(w http.ResponseWriter, r *http.Request) {
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
}
