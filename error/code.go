package error

import (
	"fmt"
	"strings"
)

// ErrorType represents the category of the error: user input, business rule, system failure, or third-party dependency.
type ErrorType string

const (
	ErrorTypeUser       ErrorType = "user"        // caused by invalid user input or action
	ErrorTypeBiz        ErrorType = "biz"         // business logic rule violation
	ErrorTypeSystem     ErrorType = "system"      // internal system fault (panic, infra, db, etc.)
	ErrorTypeThirdParty ErrorType = "third-party" // third party service / dependency issue
)

type ErrorCode struct {
	Type        ErrorType
	Application string
	Service     string
	Module      string
	Code        string
	Title       string // short summary suitable for UI heading or logs
	HTTPStatus  int
}

func NewErrorCode(errType ErrorType, application, service, module, code, title string, httpStatus int) ErrorCode {
	return ErrorCode{
		Type:        errType,
		Application: application,
		Service:     service,
		Module:      module,
		Code:        code,
		Title:       title,
		HTTPStatus:  httpStatus,
	}
}

func (e *ErrorCode) CodeString() string {
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
