package types

type ResponseWrapper struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
