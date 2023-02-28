package failures

import "fmt"

var ErrNotFound *NotFoundError
var ErrInvalidCredentials *InvalidCredentialsError

func init() {
	ErrNotFound = &NotFoundError{message: "the request resource not found"}
	ErrInvalidCredentials = &InvalidCredentialsError{message: "the credentials provided are invalid"}
}

type InvalidCredentialsError struct {
	message string
	err     error
}

func (e *InvalidCredentialsError) Error() string {
	return fmt.Sprintf("credentials error: %s", e.message)
}

func (e *InvalidCredentialsError) Unwrap() error {
	return e.err
}

type NotFoundError struct {
	message string
	err     error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.message)
}

func (e *NotFoundError) Unwrap() error {
	return e.err
}
