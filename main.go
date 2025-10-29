package main

import (
	"errors"
	"fmt"
	"lion-golang/apperr"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

	var returnedError error

	newError := apperr.New(apperr.UserNotFound, "user not found error example")
	if newError != nil {
		fmt.Println("Error occurred:", newError.Error())
		returnedError = apperr.Wrap(apperr.ResourceNotFound, "wrapping error example", newError)
	}
	if returnedError != nil {
		fmt.Println("Returned Error:", returnedError.Error())
	}
	standardError := errors.New("a standard error")
	if appErr, ok := apperr.As(returnedError); ok {
		fmt.Println("AppError Code:", appErr.Code.CodeString())
	}

	// errors.as
	if appErr, ok := apperr.As(standardError); ok {
		fmt.Println("AppError Code:", appErr.Code.CodeString())
	} else {
		fmt.Println("standardError is not an AppError")
	}

	// errors.is
	if errors.Is(returnedError, newError) {
		fmt.Println("returnedError is an AppError")
	}

	fmt.Println(apperr.IsType(returnedError, apperr.ErrorTypeUser))
	fmt.Println(apperr.IsType(returnedError, apperr.ErrorTypeBiz))

	messageError := NotFound.WithMessage("XXX not found")
	if messageError != nil {
		fmt.Println("Message Error:", messageError.Error())
	}
}

const (
	ApplicationName = "hive"
	ServiceName     = "orchestration"
)

var (
	CommonBadRequestErrorCode         = newUserErrorCode("common", "bad_request", "Bad Request", http.StatusBadRequest)
	CommonNotFoundErrorCode           = newUserErrorCode("common", "not_found", "Resource Not Found", http.StatusNotFound)
	CommonInternalErrorCode           = newSystemErrorCode("common", "internal_error", "Internal Server Error", http.StatusInternalServerError)
	CommonUnauthorizedErrorCode       = newUserErrorCode("common", "unauthorized", "Unauthorized", http.StatusUnauthorized)
	CommonForbiddenErrorCode          = newUserErrorCode("common", "forbidden", "Forbidden", http.StatusForbidden)
	CommonUnknownErrorCode            = newSystemErrorCode("common", "unknown", "Unknown Error", http.StatusInternalServerError)
	CommonServiceUnavailableErrorCode = newThirdPartyErrorCode("common", "unavailable", "Service Unavailable", http.StatusServiceUnavailable)
)

var (
	NotFound     = apperr.New(CommonNotFoundErrorCode, "resource not found")
	Internal     = apperr.New(CommonInternalErrorCode, "internal server error")
	BadRequest   = apperr.New(CommonBadRequestErrorCode, "bad request")
	Unauthorized = apperr.New(CommonUnauthorizedErrorCode, "unauthorized")
	Forbidden    = apperr.New(CommonForbiddenErrorCode, "forbidden")
	Unavailable  = apperr.New(CommonServiceUnavailableErrorCode, "service unavailable")
	Unknown      = apperr.New(CommonUnknownErrorCode, "unknown error")
)

func newUserErrorCode(module, code, title string, httpStatus int) apperr.ErrorCode {
	return apperr.NewErrorCode(apperr.ErrorTypeUser, ApplicationName, ServiceName, module, code, title, httpStatus)
}

func newBizErrorCode(module, code, title string, httpStatus int) apperr.ErrorCode {
	return apperr.NewErrorCode(apperr.ErrorTypeBiz, ApplicationName, ServiceName, module, code, title, httpStatus)
}

func newSystemErrorCode(module, code, title string, httpStatus int) apperr.ErrorCode {
	return apperr.NewErrorCode(apperr.ErrorTypeSystem, ApplicationName, ServiceName, module, code, title, httpStatus)
}

func newThirdPartyErrorCode(module, code, title string, httpStatus int) apperr.ErrorCode {
	return apperr.NewErrorCode(apperr.ErrorTypeThirdParty, ApplicationName, ServiceName, module, code, title, httpStatus)
}
