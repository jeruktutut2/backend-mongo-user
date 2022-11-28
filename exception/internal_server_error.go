package exception

type InternalServerError struct {
	Message string
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{Message: message}
}

func (err InternalServerError) Error() string {
	return "InternalServerError"
}
