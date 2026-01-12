package paymentpage

import (
	"context"
	"net/http"

	"github.com/avivbintangaringga/ayobayar/types"
)

type Handler struct {
	paymentService       types.PaymentService
	paymentMethodService types.PaymentMethodService
	processor            map[string]types.UpstreamPaymentProcessor
}

func NewHandler(paymentService types.PaymentService, paymentMethodService types.PaymentMethodService, processor map[string]types.UpstreamPaymentProcessor) *Handler {
	return &Handler{
		paymentService:       paymentService,
		paymentMethodService: paymentMethodService,
		processor:            processor,
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

	method, err := h.paymentMethodService.GetPaymentMethodById(payment.PaymentMethodId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	upstream, err := h.processor[payment.PaymentMethodId].GetPaymentResult(payment.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page(*payment, *method, upstream).Render(context.Background(), w)
}
