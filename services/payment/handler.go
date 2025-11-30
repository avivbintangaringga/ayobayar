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
		slog.Error("ListPayments", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	json.WriteSuccess(w, paymentList)
}

func (h *handler) GetPaymentDetail(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		json.WriteError(w, http.StatusBadRequest, "bad request")
		return
	}

	payment, err := h.svc.GetPaymentDetail(id)
	if err != nil {
		slog.Error("GetPaymentDetail", "error", err, "id", id)
		json.WriteError(w, http.StatusNotFound, "item not found")
		return
	}

	json.WriteSuccess(w, payment)
}

func (h *handler) PostPayment(w http.ResponseWriter, r *http.Request) {
	payment, err := h.svc.CreatePayment(types.Payment{})
	if err != nil {
		slog.Error("PostPayment", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	json.WriteSuccess(w, payment)
}
