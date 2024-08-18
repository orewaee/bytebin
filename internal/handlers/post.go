package handlers

import (
	"github.com/orewaee/bytebin/internal/app/api"
	"github.com/orewaee/bytebin/internal/app/domain"
	"github.com/orewaee/bytebin/internal/config"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"time"
)

type PostHandler struct {
	bytebinApi api.BytebinApi
	log        *zerolog.Logger
}

func NewPostHandler(bytebinApi api.BytebinApi, log *zerolog.Logger) *PostHandler {
	return &PostHandler{
		bytebinApi: bytebinApi,
		log:        log,
	}
}

func (handler *PostHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	bin, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		message := []byte("failed to read bytes")
		if _, err := writer.Write(message); err != nil {
			handler.log.Err(err).Send()
		}

		return
	}

	contentType := request.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(bin)
	}

	id := xid.New().String()

	ip := request.RemoteAddr
	userAgent := request.UserAgent()

	meta := &domain.Meta{
		Id:          id,
		ContentType: contentType,
		Ip:          ip,
		UserAgent:   userAgent,
		CreatedAt:   time.Now(),
		Lifetime:    config.Get().Lifetime,
	}

	if err := handler.bytebinApi.Add(id, bin, meta); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		message := []byte(err.Error())
		if _, err := writer.Write(message); err != nil {
			handler.log.Err(err).Send()
		}

		return
	}

	writer.WriteHeader(http.StatusCreated)

	message := []byte(id)
	if _, err := writer.Write(message); err != nil {
		handler.log.Err(err).Send()
	}
}
