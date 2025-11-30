package payment

import (
	"log/slog"
	"net/http"

	"github.com/avivbintangaringga/dompetkita/json"
	"github.com/avivbintangaringga/dompetkita/types"
)

type handler struct {
	svc types.PaymentService
}

func NewHandler(svc types.PaymentService) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) ListPayments(w http.ResponseWriter, r *http.Request) {
	paymentList, err := h.svc.GetPaymentList()
	if err != nil {
		slog.Error("ListPayments: %s", err)
		json.WriteError(w, 500, "internal server error")
		return
	}

	json.WriteSuccess(w, paymentList)
}
