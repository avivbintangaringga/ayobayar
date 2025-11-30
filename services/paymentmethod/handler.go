package paymentmethod

import (
	"log/slog"
	"net/http"

	"github.com/avivbintangaringga/dompetkita/json"
	"github.com/avivbintangaringga/dompetkita/types"
)

type handler struct {
	svc types.PaymentMethodService
}

func NewHandler(svc types.PaymentMethodService) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	paymentMethods, err := h.svc.GetPaymentMethods()
	if err != nil {
		slog.Error("GetPaymentMethods", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	json.WriteSuccess(w, paymentMethods)
}
