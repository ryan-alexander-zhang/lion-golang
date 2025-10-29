package apperr

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func New(code ErrorCode, msg string) *AppError {
	return &AppError{Code: code, Message: msg}
}

func Wrap(code ErrorCode, msg string, cause error) *AppError {
	if cause == nil {
		return New(code, msg)
	}
	return &AppError{Code: code, Message: msg, Cause: cause}
}

func (e *AppError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Message != "" {
		return e.Message
	}

	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Code.Title, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code.CodeString(), e.Code.Title)
}

// Unwrap exposes the underlying cause for errors.Is / errors.As.
func (e *AppError) Unwrap() error { return e.Cause }

func As(err error) (*AppError, bool) {
	var ae *AppError
	if errors.As(err, &ae) {
		return ae, true
	}
	return nil, false
}

// IsType reports whether err (possibly wrapped) is an AppError whose Code.Type equals target.
func IsType(err error, target ErrorType) bool {
	if err == nil {
		return false
	}
	var ae *AppError
	if errors.As(err, &ae) && ae != nil {
		return ae.Code.Type == target
	}
	return false
}

func (e *AppError) IsType(target ErrorType) bool {
	return e != nil && e.Code.Type == target
}

func (e *AppError) WithMessage(msg string) *AppError {
	if e == nil {
		return nil
	}
	return &AppError{
		Code:    e.Code,
		Message: msg,
		Cause:   e.Cause,
	}
}

func (e *AppError) WithCause(cause error) *AppError {
	if e == nil {
		return nil
	}
	return &AppError{
		Code:    e.Code,
		Message: cause.Error(),
		Cause:   cause,
	}
}
