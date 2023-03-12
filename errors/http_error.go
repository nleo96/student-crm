package errors

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	code        int
	message     string
	err         error
	description string
	cause       any
}

func NewHTTPError(code int, err error) *HTTPError {
	return &HTTPError{code: code, err: err}
}

func NewUnkownHTTPError(err error) *HTTPError {
	return NewHTTPError(
		http.StatusInternalServerError,
		err,
	).WithMessage("An internal server error has occurred while processing your request.")
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

func (e *HTTPError) WithMessage(message string) *HTTPError {
	e.message = message
	return e
}

func (e *HTTPError) WithDescription(description string) *HTTPError {
	e.description = description
	return e
}

func (e *HTTPError) WithCause(cause any) *HTTPError {
	e.cause = cause
	return e
}

func (e *HTTPError) MarshalJSON() ([]byte, error) {
	var message string

	if len(e.message) == 0 && e.err != nil {
		message = e.err.Error()
	} else {
		message = e.message
	}

	return json.Marshal(
		&struct {
			Message     string `json:"message"`
			Description string `json:"description,omitempty"`
			Cause       any    `json:"cause,omitempty"`
		}{
			message,
			e.description,
			e.cause,
		},
	)
}
