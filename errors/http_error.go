package errors

import "encoding/json"

type HTTPError struct {
	code    int
	message string
	err     error
	cause   any
}

func NewHTTPError(code int, message string, err error) *HTTPError {
	return &HTTPError{code, message, err, nil}
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

func (e *HTTPError) WithCause(cause any) *HTTPError {
	return &HTTPError{e.code, e.message, e.err, cause}
}

func (e *HTTPError) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Message string `json:"message"`
			Error   any    `json:"error,omitempty"`
		}{
			e.message,
			e.cause,
		},
	)
}
