package apperr

import "net/http"

const (
	ApplicationName = "hive"
	ServiceName     = "orchestration"
)

// user errors
var (
	FunctionNotFound = NewErrorCode(ErrorTypeUser, ApplicationName, ServiceName, "function", "not_found", "Function Not Found", http.StatusNotFound)
	UserNotFound     = NewErrorCode(ErrorTypeUser, ApplicationName, ServiceName, "user", "not_found", "User Not Found", http.StatusNotFound)
)

// biz errors
var (
	FunctionCreateConflict = NewErrorCode(ErrorTypeBiz, ApplicationName, ServiceName, "function", "create_conflict", "Create Conflict", http.StatusConflict)
)

// common errors
var (
	ResourceNotFound = NewErrorCode(ErrorTypeUser, ApplicationName, ServiceName, "XXX", "not_found", "", http.StatusNotFound)
)
