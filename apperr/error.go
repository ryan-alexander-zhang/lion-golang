package apperr

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	UserNotFound = NewAppError(ErrorTypeUser, ApplicationName, ServiceName, "user", "not_found", "User Not Found", http.StatusNotFound, "not found user error example")
)

type AppError struct {
	Type        ErrorType
	Application string
	Service     string
	Module      string
	Code        string
	Title       string // short summary suitable for UI heading or logs
	HTTPStatus  int
	Message     string
	Cause       error
}

func NewAppError(errType ErrorType, application, service, module, code, title string, httpStatus int, message string) *AppError {
	return &AppError{
		Type:        errType,
		Application: application,
		Service:     service,
		Module:      module,
		Code:        code,
		Title:       title,
		HTTPStatus:  httpStatus,
		Message:     message,
	}
}

func (e *AppError) SetMsg(msg string) *AppError {
	appError := NewAppError(e.Type, e.Application, e.Service, e.Module, e.Code, e.Title, e.HTTPStatus, e.Message)
	appError.Message = msg
	return appError
}

//func WrapError(code ErrorCode, msg string, cause error) *AppError {
//	if cause == nil {
//		return NewError(code, msg)
//	}
//	return &AppError{Code: code, Message: msg, Cause: cause}
//}

func (e *AppError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Message != "" {
		return e.Message
	}

	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Title, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.CodeString(), e.Title)
}

func (e *AppError) CodeString() string {
	typeInitial := ""
	if len(e.Type) > 0 {
		typeInitial = strings.ToUpper(string(e.Type[0]))
	}
	return fmt.Sprintf("%s.%s.%s.%s.%s",
		typeInitial,
		strings.ToUpper(e.Application),
		strings.ToUpper(e.Service),
		strings.ToUpper(e.Module),
		strings.ToUpper(e.Code),
	)
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
		return ae.Type == target
	}
	return false
}

func (e *AppError) IsType(target ErrorType) bool {
	return e != nil && e.Type == target
}
