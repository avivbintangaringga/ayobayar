package paymentpage

import (
	"context"
	"net/http"

	"github.com/avivbintangaringga/ayobayar/types"
)

type Handler struct {
	paymentService types.PaymentService
}

func NewHandler(paymentService types.PaymentService) *Handler {
	return &Handler{
		paymentService: paymentService,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	payment, err := h.paymentService.GetPaymentDetail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	page(*payment).Render(context.Background(), w)
}
