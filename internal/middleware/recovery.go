package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				cleanStack := CleanStackTrace()

				red := color.New(color.FgRed).PrintfFunc()
				red("[PANIC]: %v\n%v\n",
					r,
					cleanStack,
				)

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

func CleanStackTrace() string {
	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")
	var cleaned []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.Contains(trimmed, "/recovery") ||
			strings.Contains(trimmed, "/cors") ||
			strings.Contains(trimmed, "router.SetRouters.CustomRecovery") ||
			strings.Contains(trimmed, "CleanStackTrace") ||
			strings.Contains(trimmed, "SetRouters.CorsHandler") {
			continue
		}

		if strings.Contains(trimmed, "goex/") {
			cleaned = append(cleaned, line)
			continue
		}
	}

	return strings.Join(cleaned, "\n")
}
