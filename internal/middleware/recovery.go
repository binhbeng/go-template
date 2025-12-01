package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/binhbeng/goex/internal/api"
	"github.com/gin-gonic/gin"
)

func GetPanicLine(stack []byte) string {
	lines := bytes.Split(stack, []byte("\n"))
	for i := range lines {
		line := string(lines[i])
		if strings.Contains(line, "/runtime/")|| strings.Contains(line, "/middleware/") || strings.Contains(line, "/go/src/") {
			continue
		}

		if strings.Contains(line, ".go:") {
			return line
		}
	}
	return "unknown line"
}

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				fmt.Println("\033[31mPanic:", r, "\n", stack, "\033[0m")
				// panicLine := GetPanicLine(stack)
				// fmt.Println("\033[31mPanic:", r, "\n", panicLine, "\033[0m")
				api.HandleError(c, http.StatusInternalServerError, "Internal Server Error", fmt.Errorf("panic %s", r))
			}
		}()

		c.Next()
	}
}
