package paymentlistpage

import (
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
	var payment types.Payment
	payments, err := h.ps.GetPaymentList()
	if err != nil || len(payments) == 0 {
		payment = types.Payment{}
	}

	payment = payments[0]

	page(payment).Render(r.Context(), w)
}
