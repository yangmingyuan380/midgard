package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CROSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// 如果是option方法，则直接返回200
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else { // 否则继续向下执行
			ctx.Next()
		}
	}
}
