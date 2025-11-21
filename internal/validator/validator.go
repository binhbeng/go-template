package validator

import (
	"fmt"
	r "github.com/binhbeng/goex/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func CheckQueryParams(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		fmt.Println(err)
		r.ErrorResponse(c, 400, "Invalid params", err, nil)
		c.Abort()
		return err
	}

	return nil
}

func CheckPostParams(c *gin.Context, obj any) error {
	if err := c.ShouldBind(obj); err != nil {
		r.ErrorResponse(c, 400, "Invalid params", err, nil)
		c.Abort()
		return err
	}

	return nil
}