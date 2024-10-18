package controllers

import (
	"github.com/orewaee/bytebin/internal/utils"
	"net/http"
)

func (controller *RestController) get(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	bin, meta, err := controller.bytebinApi.GetById(id)
	if err != nil {
		utils.MustWriteString(writer, "not found", http.StatusNotFound)
		controller.logger.Error().Err(err).Send()
		return
	}

	writer.Header().Set("Content-Type", meta.ContentType)
	utils.MustWriteBytes(writer, bin, http.StatusOK)
}
