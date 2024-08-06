package http

import (
	"github.com/orewaee/bytebin/internal/app/domain"
	"github.com/orewaee/bytebin/internal/app/ports"
	"github.com/orewaee/bytebin/internal/config"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"time"
)

type PostHandler struct {
	bytebinService ports.BytebinService
	log            *zerolog.Logger
}

func NewPostHandler(bytebinService ports.BytebinService, log *zerolog.Logger) *PostHandler {
	return &PostHandler{
		bytebinService: bytebinService,
		log:            log,
	}
}

func (handler *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.log.Debug().
		Str("method", "POST").
		Str("route", r.URL.String()).
		Send()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	bin, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		message := []byte("failed to read bytes")
		if _, err := w.Write(message); err != nil {
			handler.log.Err(err).Send()
		}

		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(bin)
	}

	id := xid.New().String()

	ip := r.RemoteAddr
	userAgent := r.UserAgent()

	meta := &domain.Meta{
		Id:          id,
		ContentType: contentType,
		Ip:          ip,
		UserAgent:   userAgent,
		CreatedAt:   time.Now(),
		Lifetime:    config.Get().Lifetime,
	}

	if err := handler.bytebinService.Add(id, bin, meta); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		message := []byte(err.Error())
		if _, err := w.Write(message); err != nil {
			handler.log.Err(err).Send()
		}

		return
	}

	w.WriteHeader(http.StatusCreated)

	message := []byte(id)
	if _, err := w.Write(message); err != nil {
		handler.log.Err(err).Send()
	}
}
