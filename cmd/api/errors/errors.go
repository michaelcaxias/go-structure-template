package errors

import "net/http"

type APIError struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Cause   []map[string]string `json:"cause"`
}

func NewAPIError(status int, message string, cause []map[string]string) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
		Cause:   cause,
	}
}

func NewBadRequestError(message string, cause []map[string]string) *APIError {
	return NewAPIError(http.StatusBadRequest, message, cause)
}
