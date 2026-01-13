package api

import "github.com/gin-gonic/gin"

type Response struct {
	Status  bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func HandleSuccess(ctx *gin.Context, statusCode int, message string, data ...any) {
	response := Response{
		Status:  true,
		Message: message,
		Error:   nil,
		Data:    data,
	}

	if len(data) > 0 && data[0] != nil {
		if m, ok := data[0].(map[string]any); ok {
			// if p, exists := m["pagination"]; exists {
			// 	response.Pagination = p
			// }

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

func HandleError(ctx *gin.Context, statusCode int, message string, err error) {
	response := Response{
		Status:  false,
		Message: message,
		Error:   err.Error(),
	}
	// ctx.Error(fmt.Errorf("%s",  err.Error()))
	ctx.AbortWithStatusJSON(statusCode, response)
}