package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)

func SetLogLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

func NewLogger(env string) (*zap.Logger, error) {
	if env == "" || env == DevEnv {
		return zap.NewDevelopment()
	}
	return NewProductionConfig().Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

// NewProductionConfig builds a reasonable default production logging
// configuration.
// Logging is enabled at InfoLevel and above, and uses a JSON encoder.
// Logs are written to standard error.
// Stacktraces are included on logs of ErrorLevel and above.
// DPanicLevel logs will not panic, but will write a stacktrace.
//
// Sampling is enabled at 100:100 by default,
// meaning that after the first 100 log entries
// with the same level and message in the same second,
// it will log every 100th entry
// with the same level and message in the same second.
// You may disable this behavior by setting Sampling to nil.
//
// See [NewProductionEncoderConfig] for information
// on the default encoder configuration.
func NewProductionConfig() zap.Config {
	return zap.Config{
		Level:       atomicLevel,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// NewProductionEncoderConfig returns an opinionated EncoderConfig for
// production environments.
//
// Messages encoded with this configuration will be JSON-formatted
// and will have the following keys by default:
//
//   - "level": The logging level (e.g. "info", "error").
//   - "ts": The current time in number of seconds since the Unix epoch.
//   - "msg": The message passed to the log statement.
//   - "caller": If available, a short path to the file and line number
//     where the log statement was issued.
//     The logger configuration determines whether this field is captured.
//   - "stacktrace": If available, a stack trace from the line
//     where the log statement was issued.
//     The logger configuration determines whether this field is captured.
//
// By default, the following formats are used for different types:
//
//   - Time is formatted as floating-point number of seconds since the Unix
//     epoch.
//   - Duration is formatted as floating-point number of seconds.
//
// You may change these by setting the appropriate fields in the returned
// object.
// For example, use the following to change the time encoding format:
//
//	cfg := zap.NewProductionEncoderConfig()
//	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
func NewProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     utcISO8601TimeEncoder, // changed to explicit UTC ISO8601
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// utcISO8601TimeEncoder encodes time in UTC using RFC3339Nano with trailing 'Z'.
func utcISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(time.RFC3339Nano))
}
