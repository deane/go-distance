package api

import (
	"encoding/json"
	"fmt"
)

// errorResponse represents the Json object
// returned in case of an error
type errorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func jsonError(code int, message string) []byte {
	errStruct := errorResponse{code, message}
	response, err := json.Marshal(errStruct)
	fmt.Println(err)
	return response
}
