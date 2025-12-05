package common

import (
	"log/slog"
	"net/http"

	"github.com/avivbintangaringga/ayobayar/json"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	slog.Error("resource not found", "url", r.URL.Path)
	json.WriteError(w, http.StatusNotFound, "resource not found")
}

func (h *handler) HandleNotAllowed(w http.ResponseWriter, r *http.Request) {
	slog.Error("method not allowed", "url", r.URL.Path, "method", r.Method)
	json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
}
