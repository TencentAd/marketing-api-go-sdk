package errors

// ResponseError ...
type ResponseStrictError struct {
	Errors string
}

// Error ...
func (e ResponseStrictError) Error() string {
	return e.Errors
}

// NewResponseStrictError ...
func NewResponseStrictError(errors string) ResponseStrictError {
	return ResponseStrictError{
		Errors: errors,
	}
}
