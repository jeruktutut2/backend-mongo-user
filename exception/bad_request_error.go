package exception

type BadRequestError struct {
	Message string
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{Message: message}
}

func (err BadRequestError) Error() string {
	return "BadRequestError"
}
