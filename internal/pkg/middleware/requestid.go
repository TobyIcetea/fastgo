package middleware

import (
	"github.com/TobyIcetea/fastgo/internal/pkg/contextx"
	"github.com/TobyIcetea/fastgo/internal/pkg/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIUD 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context，response 中注入 `x-request-id` 键值对
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 `x-request-id`，如果不存在则生成新的 UUID
		requestID := c.Request.Header.Get(known.XRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将 RequestID 保存到 context.Context 中，以便后续程序使用
		ctx := contextx.WithRequestID(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)

		// 将 RequestID 保存到 HTTP 请求头中，Header 的键为 `x-request-id`
		c.Writer.Header().Set(known.XRequestID, requestID)

		// 继续处理请求
		c.Next()
	}
}
