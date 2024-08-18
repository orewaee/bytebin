package handlers

import (
	"github.com/orewaee/bytebin/internal/app/api"
	"github.com/orewaee/bytebin/internal/utils"
	"github.com/rs/zerolog"
	"net/http"
)

type GetHandler struct {
	bytebinApi api.BytebinApi
	log        *zerolog.Logger
}

func NewGetHandler(bytebinApi api.BytebinApi, log *zerolog.Logger) *GetHandler {
	return &GetHandler{
		bytebinApi: bytebinApi,
		log:        log,
	}
}

func (handler *GetHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	bin, meta, err := handler.bytebinApi.GetById(id)
	if err != nil {
		utils.MustWriteString(writer, "not found", http.StatusNotFound)
		handler.log.Error().Err(err).Send()
		return
	}

	writer.Header().Set("Content-Type", meta.ContentType)
	utils.MustWriteBytes(writer, bin, http.StatusOK)
}
