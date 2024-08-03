package http

import (
	"github.com/orewaee/bytebin/internal/app/ports"
	"log"
	"net/http"
)

type GetHandler struct {
	bytebinService ports.BytebinService
}

func NewGetHandler(bytebinService ports.BytebinService) *GetHandler {
	return &GetHandler{
		bytebinService: bytebinService,
	}
}

func (handler *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	bin, meta, err := handler.bytebinService.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		message := []byte("not found")
		if _, err := w.Write(message); err != nil {
			log.Println(err)
		}

		return
	}

	w.Header().Set("Content-Type", meta.ContentType)
	if _, err := w.Write(bin); err != nil {
		log.Println(err)
	}
}
