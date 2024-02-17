package errors

import (
	"fmt"
	"net/http"
)

type ErrorType string

const (
	ErrInternalServerError ErrorType = "Internal Server Error"
	ErrNotFound            ErrorType = "Your requested Item is not found"
	ErrConflict            ErrorType = "Your Item already exist"
	ErrBadRequest          ErrorType = "Bad Request"
	ErrorMaxLimit          ErrorType = "Max limit reached"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(errType ErrorType, err error) *ErrorResponse {
	return &ErrorResponse{
		Message: fmt.Sprintf("%s: %s", errType, err.Error()),
	}
}

func NewInternalServerError(err error) (int, *ErrorResponse) {
	return http.StatusInternalServerError, NewErrorResponse(ErrInternalServerError, err)
}

func NewNotFoundError(err error) (int, *ErrorResponse) {
	return http.StatusNotFound, NewErrorResponse(ErrNotFound, err)
}

func NewConflictError(err error) (int, *ErrorResponse) {
	return http.StatusConflict, NewErrorResponse(ErrConflict, err)
}

func NewBadRequestError(err error) (int, *ErrorResponse) {
	return http.StatusBadRequest, NewErrorResponse(ErrBadRequest, err)
}

func NewMaxLimitError() (int, *ErrorResponse) {
	err := fmt.Errorf("max limit reached")
	return http.StatusBadRequest, NewErrorResponse(ErrorMaxLimit, err)
}
