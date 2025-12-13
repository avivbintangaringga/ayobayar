package paymentlistpage

import (
	"log/slog"
	"net/http"

	"github.com/avivbintangaringga/ayobayar/types"
)

type Handler struct {
	ps types.PaymentService
}

func NewHandler(ps types.PaymentService) *Handler {
	return &Handler{
		ps: ps,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	payments, err := h.ps.GetPaymentList()
	if err != nil || len(payments) == 0 {
		slog.Error("PAYMENTLISTPAGE", "error", err)
	}

	page(payments).Render(r.Context(), w)
}
