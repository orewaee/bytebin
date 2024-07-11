package handlers

import "net/http"

type GetHandler struct{}

func NewGetHandler() *GetHandler {
	return &GetHandler{}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
