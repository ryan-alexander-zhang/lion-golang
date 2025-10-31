package main

import (
	"errors"
	error2 "lion-golang/error"
	"lion-golang/logger"
	"net/http"

	"go.uber.org/zap"
)

func LoggerDemo() {
	appLogger, err := logger.NewLogger(logger.DevEnv)
	if err != nil {
		panic(err)
	}
	appLogger.Sugar().Info("Logger Demo")
	appleLogger := appLogger.With(zap.String("component", "apple"))
	appleLogger.Sugar().Info("Logger Apple")

	// Log error example
	err = errors.New("Error Demo")
	//appleLogger.Sugar().Errorf("An error occurred: %v", err)

	appError := error2.New(error2.NewErrorCode(error2.ErrorTypeBiz, "app", "svc", "test", "bad", "Bad request", http.StatusBadRequest), "An error occurred: %v")

	appleLogger.Sugar().Errorw("A wrapped error occurred", "error", err, "cause", appError.Unwrap())

	// Wrap error example
	//wrappedErr := error.Wrap(error.FunctionNotFound, "new message", err)

	//appleLogger.Sugar().Errorw("A wrapped error occurred", "error", wrappedErr, "cause", wrappedErr.Unwrap())
}
