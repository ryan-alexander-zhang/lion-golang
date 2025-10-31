package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"lion-golang/logger"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	shutdown := initOTel()
	defer shutdown(context.Background())

	// New appLogger
	appLogger, err := logger.NewLogger(logger.ProdEnv)
	if err != nil {
		panic(err)
	}

	// Create a Gin router with no default middleware (appLogger and recovery)
	r := gin.New()

	// OpenTelemetry middleware
	r.Use(otelgin.Middleware("your-service-name"))

	// Add a request ID middleware
	// Fetch or generate a request ID for each request and set it in the context
	// Accessible via requestid.Get(c)
	// Also sets "X-Request-Id" header in the response
	r.Use(requestid.New())

	// Middleware to add logger with request ID, trace and span id field to context
	r.Use(addTraceMiddleware(appLogger))
	r.Use(addTraceResponseMiddleware())
	r.Use(logRequestMiddleware(appLogger))
	r.Use(errorHandlerMiddleware())

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(appLogger, true))

	// =========================================================================
	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		appLogger.Sugar().Info("ping coming")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		// Return JSON response
		appLogger.Sugar().Info("test coming")
		service := NewService(appLogger)
		service.DoSomething()
		c.JSON(http.StatusOK, gin.H{
			"message": "on",
		})
	})

	r.GET("/error", func(c *gin.Context) {
		// Return JSON response
		c.Error(errors.New("error"))
		return
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

// Add a ginzap middleware, which:
//   - Logs all requests, like a combined access and error log.
//   - Logs to stdout.
//   - RFC3339 with UTC time format.
func logRequestMiddleware(appLogger *zap.Logger) gin.HandlerFunc {
	return ginzap.GinzapWithConfig(appLogger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
			fields := []zapcore.Field{}
			// log request ID
			if requestID := requestid.Get(c); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}

			// log trace and span ID
			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
				fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
			}

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))

			return fields
		}),
		Skipper: func(c *gin.Context) bool {
			return c.Request.URL.Path == "/healthz" && c.Request.Method == "GET"
		},
	})
}

// Middleware to set tracing headers in the response
// - X-Request-Id
// - traceparent
// - Server-Timing
func addTraceResponseMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// X-Request-ID
		if rid := requestid.Get(c); rid != "" {
			c.Header("X-Request-Id", rid)
		}

		// traceparent + Server-Timing
		sc := trace.SpanFromContext(c.Request.Context()).SpanContext()
		if sc.IsValid() {
			// W3C traceparent
			tp := "00-" + sc.TraceID().String() + "-" + sc.SpanID().String() + "-01"
			c.Header("traceparent", tp)
			c.Header("Server-Timing", `traceparent;desc="`+tp+`"`)
		}
		c.Next()
	}
}

func addTraceMiddleware(appLogger *zap.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Add request ID to logger
		reqID := requestid.Get(c)
		appLogger = appLogger.With(zap.String("request_id", reqID))

		// Add trace_id and span_id to logger
		sc := trace.SpanFromContext(c.Request.Context()).SpanContext()
		if sc.IsValid() {
			appLogger = appLogger.With(
				zap.String("trace_id", sc.TraceID().String()),
				zap.String("span_id", sc.SpanID().String()),
			)
		}
		c.Set("logger", appLogger.With(zap.String("request_id", reqID)))
		c.Next()
	}
}

func initOTel() func(context.Context) error {
	// TODO : configure OTLP exporter to send data to Collector
	// For demo, we use stdout exporter
	// 1) Exporter: use stdout for demo; for production prefer OTLP -> Collector
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	// 2) Resource: set service name
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(semconv.ServiceNameKey.String("your-service-name")),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 3) TracerProvider: batch processor + resource
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
	)

	// 4) Set global TracerProvider and TextMap Propagator
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	return tp.Shutdown
}

// ErrorHandler captures errors and returns a consistent JSON error response
func errorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a generic error message
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":    false,
				"errCode":    "InternalServerError",
				"errMessage": err.Error(),
			})
		}

		// Any other steps if no errors are found
	}
}
