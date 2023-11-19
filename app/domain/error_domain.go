package domain

type Error struct {
	message string
}

func NewError(message string) *Error {
	return &Error{
		message: message,
	}
}

func (e *Error) Error() string {
	return e.message
}
