package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v\n", r)
				fmt.Printf("%s\n", debug.Stack())

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal Server Error",
					"error":   fmt.Sprintf("Panic: %v", r),
				})
			}
		}()

		c.Next()
	}
}
