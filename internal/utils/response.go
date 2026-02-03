package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status     bool   `json:"success"`
	Message    string `json:"message"`
	Error      any    `json:"error,omitempty"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data ...any) {
	response := Response{
		Status:  true,
		Message: message,
		Data:    data,
		Error:   nil,
	}

	if len(data) > 0 && data[0] != nil {
		if m, ok := data[0].(map[string]any); ok {
			if p, exists := m["pagination"]; exists {
				response.Pagination = p
			}

			if d, exists := m["data"]; exists {
				response.Data = d
			} else {
				response.Data = m
			}
		} else {
			response.Data = data[0]
		}
	}

	ctx.AbortWithStatusJSON(statusCode, response)
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, e any) {
	if len(message) > 0 {
		message = strings.ToUpper(string(message[0])) + message[1:]
	}

	var err string

	switch e := e.(type) {
	case nil:
		err = ""

	case map[string]string:
		var parts []string
		for _, msg := range e {
			parts = append(parts, fmt.Sprintf("%s", msg))
		}
		err = strings.Join(parts, ", ")

	case error:
		err = e.Error()

	default:
		err = fmt.Sprintf("%v", e)
	}

	if len(err) > 0 {
		err = strings.ToUpper(string(err[0])) + err[1:]
	}

	response := Response{
		Status:  false,
		Message: message,
		Error:   err,
	}

	ctx.AbortWithStatusJSON(statusCode, response)
}

func HttpBadRequest(ctx *gin.Context, message string, e any) {
	ErrorResponse(ctx, http.StatusBadRequest, message, e)
}

func HttpNotFound(ctx *gin.Context, message string, e any) {
	ErrorResponse(ctx, http.StatusNotFound, message, e)
}

func HttpUnauthorized(ctx *gin.Context, message string, e any) {
	ErrorResponse(ctx, http.StatusUnauthorized, message, e)
}

func HttpForbidden(ctx *gin.Context, message string, e any) {
	ErrorResponse(ctx, http.StatusForbidden, message, e)
}
