package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Response struct {
	Status  bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func HandleSuccess(ctx *gin.Context, statusCode int, message string, data ...any) {
	// log.Printf("\033[0;32m%s\033[0m\n", message)
	response := Response{
		Status:  true,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	ctx.AbortWithStatusJSON(statusCode, response)
}

func HandleError(ctx *gin.Context, statusCode int, message string, err error, data any) {
	// log.Printf("\033[0;31m%s\033[0m\n", err.Error())
	errFields := strings.Split(err.Error(), "\n")
	response := Response{
		Status:  false,
		Message: message,
		Error:   errFields,
		Data:    data,
	}
	ctx.Error(fmt.Errorf("%s",  errFields))
	ctx.AbortWithStatusJSON(statusCode, response)
}

func CheckQueryParams(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		fmt.Println(err)
		HandleSuccess(c, 400, "Invalid params", err, nil)
		c.Abort()
		return err
	}

	return nil
}

func CheckPostParams(c *gin.Context, obj any) error {
	if err := c.ShouldBind(obj); err != nil {
		HandleError(c, 400, "Invalid params", err, nil)
		c.Abort()
		return err
	}

	return nil
}