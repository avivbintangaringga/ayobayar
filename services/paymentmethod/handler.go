package paymentmethod

import (
	"log/slog"
	"net/http"

	"github.com/avivbintangaringga/ayobayar/json"
	"github.com/avivbintangaringga/ayobayar/types"
)

type Handler struct {
	svc types.PaymentMethodService
}

func NewHandler(svc types.PaymentMethodService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	paymentMethods, err := h.svc.GetPaymentMethods()
	if err != nil {
		slog.Error("GetPaymentMethods", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	json.WriteSuccess(w, paymentMethods)
}
