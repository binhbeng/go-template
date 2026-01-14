package api

import "github.com/gin-gonic/gin"

func GetUserIdFromCtx(c *gin.Context) int64 {
	return c.GetInt64("user_id")
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