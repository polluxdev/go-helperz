package helperz

import "context"

type CtxHelper interface {
	WithRequestID(ctx context.Context, requestID string) context.Context
	GetRequestID(ctx context.Context) string
}

type contextKey string

const REQUEST_ID_KEY contextKey = "requestID"

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, REQUEST_ID_KEY, requestID)
}

func GetRequestID(ctx context.Context) string {
	if val, ok := ctx.Value(REQUEST_ID_KEY).(string); ok {
		return val
	}
	return ""
}
