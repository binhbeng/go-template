package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(500, gin.H{
			"error": fmt.Sprintf("%v", "PanicExceptionRecord"),
		})
	}
}
