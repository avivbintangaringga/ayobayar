package payment

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/avivbintangaringga/dompetkita/json"
	"github.com/avivbintangaringga/dompetkita/types"
)

type Handler struct {
	svc types.PaymentService
}

func NewHandler(svc types.PaymentService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) ListPayments(w http.ResponseWriter, r *http.Request) {
	paymentList, err := h.svc.GetPaymentList()
	if err != nil {
		slog.Error("ListPayments", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	json.WriteSuccess(w, paymentList)
}

func (h *Handler) GetPaymentDetail(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) PostPayment(w http.ResponseWriter, r *http.Request) {
	var data types.PaymentRequest
	if err := json.ReadRequestBody(r, &data); err != nil {
		slog.Error("PostPayment", "json parse", err)
		json.WriteError(w, http.StatusBadRequest, "bad request")
		return
	}

	if err := json.Validate(data); err != nil {
		slog.Error("PostPayment", "validation", err)
		json.WriteError(w, http.StatusBadRequest, "bad request")
		return
	}

	paymentData := types.Payment{
		Amount:          data.Amount,
		Desc:            data.Desc,
		CallbackUrl:     data.CallbackUrl,
		MerchantId:      data.MerchantId,
		MerchantOrderId: data.MerchantOrderId,
		UserEmail:       data.UserEmail,
		UserName:        data.UserName,
		PaymentMethodId: data.PaymentMethodId,
	}

	payment, err := h.svc.CreatePayment(paymentData)
	if err != nil {
		slog.Error("PostPayment", "error", err)

		if errors.Is(err, types.ErrValidation) {
			json.WriteError(w, http.StatusBadRequest, "bad request")
			return
		}

		json.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	json.WriteSuccess(w, payment)
}
