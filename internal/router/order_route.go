package router

import (
	"github.com/binhbeng/goex/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetOrderApiRoute(api *gin.RouterGroup, orderHandler *handler.OrderHandler) {
	{
		api.GET("orders", orderHandler.GetListOrder)
		api.POST("order", orderHandler.CreateOrder)
	}
}
