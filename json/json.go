package json

import (
	"encoding/json"
	"net/http"

	"github.com/avivbintangaringga/dompetkita/types"
)

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
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
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
