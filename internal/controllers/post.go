package controllers

import (
	"github.com/orewaee/bytebin/internal/app/domain"
	"github.com/orewaee/bytebin/internal/config"
	"github.com/orewaee/bytebin/internal/utils"
	"github.com/rs/xid"
	"io"
	"net/http"
	"time"
)

func (controller *RestController) post(writer http.ResponseWriter, request *http.Request) {
	bin, err := io.ReadAll(request.Body)
	if err != nil {
		utils.MustWriteString(writer, "failed to read bytes", http.StatusInternalServerError)
		controller.logger.Error().Err(err).Send()
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

	if err := controller.bytebinApi.Add(id, bin, meta); err != nil {
		utils.MustWriteString(writer, err.Error(), http.StatusInternalServerError)
		controller.logger.Error().Err(err).Send()
		return
	}

	utils.MustWriteString(writer, id, http.StatusCreated)
}
