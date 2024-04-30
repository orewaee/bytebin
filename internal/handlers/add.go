package handlers

import (
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/rs/xid"
	"io"
	"net/http"
	"time"
)

type AddHandler struct {
	Limit    int
	Lifetime time.Duration
}

func NewAddHandler(limit int, lifetime time.Duration) *AddHandler {
	return &AddHandler{
		Limit:    limit,
		Lifetime: lifetime,
	}
}

func (handler *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(bytes) > handler.Limit {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")

	id := xid.New().String()

	storage.AddBin(id, &bin.Bin{
		ContentType: contentType,
		Bytes:       bytes,
	})

	go func() {
		time.Sleep(handler.Lifetime)
		storage.RemoveBin(id)
	}()

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
