package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/binhbeng/goex/config"
	masker "github.com/coopnorge/go-masker-lib"
	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

var (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[97;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		bodyReq:=""

		if config.C.App.EnableBodyLog == true {
			blw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			c.Writer = blw
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}
	
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			bodyReq= MaskAndLogJSON(bodyBytes)
		}

		c.Next()

		now := time.Now().Format("2006/01/02 15:04:05")
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		ip := c.ClientIP()
		fullPath := c.Request.URL.RequestURI()
		errors := c.Errors.String()
		
		fmt.Printf("%-7s %s | %s%-3d%s | %s | %15s | %s%-7s%s %-25s | %s \n",
			"[GOEX]",
			now,
			colorForStatus(status), status, reset,
			latency,
			ip,
			colorForMethod(method), method, reset,
			fullPath,
			bodyReq,
		)

		if errors != "" {
			log.Printf("\033[0;31m%s\033[0m", strings.TrimSpace(errors))
		}

	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}

func MaskAndLogJSON(bodyBytes []byte) string {
	sensitiveFields := []string{
		"password", "api_key", "token",
	}

	var obj map[string]any
	if err := json.Unmarshal(bodyBytes, &obj); err != nil {
		return string(bodyBytes)
	}

	for _, field := range sensitiveFields {
		if val, ok := obj[field]; ok {
			obj[field] = masker.CensoredString(fmt.Sprint(val))
		}
	}

	maskedB, err := json.Marshal(obj)
	if err != nil {
		return string(bodyBytes)
	}

	return string(maskedB)
}
