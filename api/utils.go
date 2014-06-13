package api

import (
	"encoding/json"
)

// The JSON object returned in the case of an API
// error.
type errorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

// jsonError returns the response to be sent given the error code
// and the error message.
func jsonError(code int, message string) []byte {
	errStruct := errorResponse{code, message}
	response, _ := json.Marshal(errStruct)
	return response
}
