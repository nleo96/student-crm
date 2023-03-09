package errors

type HTTPError struct {
	code    int
	message string
	err     error
}

func NewHTTPError(code int, message string, err error) *HTTPError {
	return &HTTPError{code, message, err}
}

func (e *HTTPError) Code() int {
	return e.code
}

func (e *HTTPError) Error() string {
	return e.message
}

func (e *HTTPError) Wrap(err error) error {
	return nil
}

func (e *HTTPError) Unwrap() error {
	return e.err
}
