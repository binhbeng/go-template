package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func GetUserIdFromCtx(c *gin.Context) uint {
    return c.GetUint("user_id")
}
