package logx

import (
	"context"
	"log"
)

type contextKey string

var loggerContextKey contextKey = "__logger"

//WithLogger ties loggers to contexts
func WithLogger(ctx context.Context, l *log.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, l)
}

//Default returns a standard logger tied to the context used by the request-level output function.
//If no logger is tied to the context, the package-level logger is returned.
func Default(ctx context.Context) *log.Logger {
	if l, ok := ctx.Value(loggerContextKey).(*log.Logger); ok && l != nil {
		return l
	}
	return log.Default()
}
