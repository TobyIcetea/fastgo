package contextx

import (
	"context"
)

// 定义用于上下文的键
type (
	// requestIDKey 定义请求 ID 的上下文键
	requestIDKey struct{}
)

// WithRequestID 将请求 ID 存放到上下文中
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// RequestID 从上下文中提取请求 ID
func RequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}
