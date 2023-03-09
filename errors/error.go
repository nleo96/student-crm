package errors

import (
	"errors"
)

type ErrorKind uint

const (
	ErrKindDatabase = (iota + 1) << 16
	ErrKindDataNotFound
	ErrKindDataDuplicated
	ErrKindInvalidCredentials
	ErrKindHashing
	ErrKindTokenSigning
	ErrKindDataBind
)

var (
	ErrDatabase           = NewError(ErrKindDatabase, "invalid database operation")
	ErrDataNotFound       = NewError(ErrKindDataNotFound, "invalid database operation")
	ErrDataDuplicated     = NewError(ErrKindDataDuplicated, "invalid database operation")
	ErrInvalidCredentials = NewError(ErrKindInvalidCredentials, "invalid credentials")
	ErrHashing            = NewError(ErrKindHashing, "invalid credentials")
	ErrTokenSigning       = NewError(ErrKindTokenSigning, "hash generation error")
	ErrDataBind           = NewError(ErrKindDataBind, "")
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

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Wrap(err error) error {
	return nil
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
