package api

import (
	"github.com/gin-gonic/gin"
)

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

func CheckQueryParams(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		HandleError(c, 400, "Invalid params", err)
		c.Abort()
		return err
	}

	return nil
}

func CheckPostParams(c *gin.Context, obj any) error {
	if err := c.ShouldBind(obj); err != nil {
		HandleError(c, 400, "Invalid params", err)
		c.Abort()
		return err
	}

	return nil
}