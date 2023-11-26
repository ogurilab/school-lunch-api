package errors

const (
	ErrInternalServerError = "Internal Server Error"
	ErrNotFound            = "Your requested Item is not found"
	ErrConflict            = "Your Item already exist"
	ErrBadParamInput       = "Given Param is not valid"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}

func NewInternalServerErrorResponse() *ErrorResponse {
	return NewErrorResponse(ErrInternalServerError)
}

func NewNotFoundErrorResponse() *ErrorResponse {
	return NewErrorResponse(ErrNotFound)
}

func NewConflictErrorResponse() *ErrorResponse {
	return NewErrorResponse(ErrConflict)
}

func NewBadParamInputErrorResponse() *ErrorResponse {
	return NewErrorResponse(ErrBadParamInput)
}
