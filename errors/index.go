package errors

import (
	"errors"
)

type ErrorKind uint

const (
	ErrKindDatabase = (iota + 1) << 32
	ErrKindDataNotFound
	ErrKindDataDuplicated
	ErrKindInvalidCredentials
	ErrKindHashing
	ErrKindTokenSigning
	ErrKindDataBind
	ErrInvalidOptionKind
)

var (
	ErrDatabase           = NewError(ErrKindDatabase, "There was an error processing your request due to a database issue")
	ErrDataNotFound       = NewError(ErrKindDataNotFound, "The requested data could not be found. Please check your request parameters and try again.")
	ErrDataDuplicated     = NewError(ErrKindDataDuplicated, "The record you're trying to add already exists in the system. Please check your request parameters and try again")
	ErrInvalidCredentials = NewError(ErrKindInvalidCredentials, "Your credentials are invalid. Please check your username and password and try again")
	ErrHashing            = NewError(ErrKindHashing, "There was an issue processing your request due to an error in hashing data. ")
	ErrTokenSigning       = NewError(ErrKindTokenSigning, "There was an issue generating your authentication token. Please try again later.")
	ErrDataBind           = NewError(ErrKindDataBind, "There was an issue processing your request due to an error in binding data. Please check the request data and ensure it is correctly formatted.")
	ErrInvalidOption      = NewError(ErrInvalidOptionKind, "Invalid option passed to function context. Try: ScanModeSoft or ScanModeDeep from this package.")
)

type Error struct {
	kind    ErrorKind
	message string
	err     error
}

func NewError(kind ErrorKind, message string) *Error {
	return &Error{kind, message, nil}
}

func (e *Error) Kind() ErrorKind {
	return e.kind
}

func (e Error) Error() string {
	return e.message
}

func (e *Error) Wrap(err error) error {
	return &Error{
		kind:    e.kind,
		message: e.message,
		err:     err,
	}
}

func (e *Error) Unwrap() error {
	return e.err
}

func New(text string) error {
	return errors.New(text)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}
