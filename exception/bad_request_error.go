package exception

type BadRequestError struct {
	Message interface{}
}

func NewBadRequestError(message interface{}) BadRequestError {
	return BadRequestError{Message: message}
}

func (err BadRequestError) Error() string {
	return "BadRequestError"
}
