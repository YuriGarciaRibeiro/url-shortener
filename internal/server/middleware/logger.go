package middleware

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

func Logger(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        logger.Info("Request",
            zap.String("path", c.Request.URL.Path),
            zap.String("method", c.Request.Method),
        )
        c.Next()
    }
}