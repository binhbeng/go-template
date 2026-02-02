package router

import (
	"github.com/binhbeng/goex/internal/app"
	"github.com/gin-gonic/gin"
)

func SetOrderApiRoute(api *gin.RouterGroup) {
	orderHandler := app.NewOrderModule().Handler()
	{
		api.GET("orders", orderHandler.GetListOrder)
	}
}
