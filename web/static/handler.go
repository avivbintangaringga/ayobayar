package static

import (
	"io/fs"
	"log/slog"
	"net/http"
)

type Handler struct {
	fs     fs.FS
	prefix string
}

func NewHandler(fs fs.FS, prefix string) *Handler {
	return &Handler{
		fs:     fs,
		prefix: prefix,
	}
}

func (h *Handler) ServeStatic(w http.ResponseWriter, r *http.Request) {
	slog.Info("FS", "fs", h.fs)
	// root := http.Dir("static")
	fileServer := http.FileServer(http.FS(h.fs))
	http.StripPrefix(h.prefix, fileServer).ServeHTTP(w, r)
}
