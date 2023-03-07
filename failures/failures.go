package failures

import (
	"fmt"
)

var (
	ErrDataNotFound = NewInternalError("register not found")
	ErrDatabaseOperation = NewInternalError("invalid database operation")
)

var ErrNotFound *NotFoundError
var ErrInvalidCredentials *InvalidCredentialsError
var ErrPasswordHashGeneration *PasswordHashGenerationError

func init() {
	ErrNotFound = &NotFoundError{
		message: "the request resource not found",
	}
	ErrInvalidCredentials = &InvalidCredentialsError{
		message: "the credentials provided are invalid",
	}
	ErrPasswordHashGeneration = &PasswordHashGenerationError{
		message: "",
	}
}

type Error interface {
	Wrap(err error) error
	Unwrap() error
	error
}

type BadRequestError struct {
	code    int
	message string
	err     error
}

func NewBadRequestError(code int, message string) *BadRequestError {
	return &BadRequestError{code, message, nil}
}

func (e *BadRequestError) Code() int {
	return e.code
}

func (e *BadRequestError) Error() string {
	return e.message
}

func (e *BadRequestError) Wrap(err error) error {
	return &BadRequestError{e.code, e.message, err}
}

func (e *BadRequestError) Unwrap() error {
	return e.err
}

type InternalError struct {
	message string
	err     error
}

func (e *InternalError) Error() string {
	return e.message
}

func (e *InternalError) Wrap(err error) error {
	return &InternalError{e.message, err}
}

func (e *InternalError) Unwrap() error {
	return e.err
}

func NewInternalError(message string) *InternalError {
	return &InternalError{message, nil}
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

type PasswordHashGenerationError struct {
	message string
	err     error
}

func (e *PasswordHashGenerationError) Error() string {
	return fmt.Sprintf("not found: %s", e.message)
}

func (e *PasswordHashGenerationError) Unwrap() error {
	return e.err
}
