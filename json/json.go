package json

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/avivbintangaringga/ayobayar/types"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func Validate(v any) error {
	err := validate.Struct(v)
	if err != nil {
		slog.Error("Validate", "error", err)
		return types.ErrValidation
	}
	return nil
}

func WriteSuccess(w http.ResponseWriter, data any) {
	WriteSuccessWithCode(w, 200, data)
}

func WriteSuccessWithCode(w http.ResponseWriter, statusCode int, data any) {
	payload := createPayload(true, statusCode, "action performed successfully!", data)
	WriteJsonResponse(w, statusCode, payload)
}

func WriteError(w http.ResponseWriter, statusCode int, msg string) {
	payload := createPayload(false, statusCode, msg, nil)
	WriteJsonResponse(w, statusCode, payload)
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func createPayload(success bool, statusCode int, message string, data any) *types.ResponseWrapper {
	return &types.ResponseWrapper{
		Success: success,
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
}

func ReadRequestBody(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ToStringWithFallback(v any, fallback string) string {
	var out []byte
	out, err := json.Marshal(v)
	if err != nil {
		out = []byte(fallback)
	}
	return string(out)
}
