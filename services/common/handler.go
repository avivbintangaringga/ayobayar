package common

import (
	"net/http"

	"github.com/avivbintangaringga/ayobayar/json"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	json.WriteError(w, http.StatusNotFound, "resource not found")
}

func (h *handler) HandleNotAllowed(w http.ResponseWriter, r *http.Request) {
	json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
}
