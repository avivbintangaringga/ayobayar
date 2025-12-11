package callback

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) DompetKita(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	w.Write([]byte("dompet kita"))
}
