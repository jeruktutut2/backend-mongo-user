package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundErro(error string) NotFoundError {
	return NotFoundError{Error: error}
}
