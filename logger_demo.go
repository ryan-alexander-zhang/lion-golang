package main

import (
	"errors"
	"lion-golang/logger"

	"go.uber.org/zap"
)

func LoggerDemo() {
	appLogger, err := logger.NewLogger(logger.ProdEnv)
	if err != nil {
		panic(err)
	}
	appLogger.Sugar().Info("Logger Demo")
	appleLogger := appLogger.With(zap.String("component", "apple"))
	appleLogger.Sugar().Info("Logger Apple")

	// Log error example
	err = errors.New("Error Demo")
	appleLogger.Sugar().Errorf("An error occurred: %v", err)
}
