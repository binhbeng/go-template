package router

import (
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserApiRoute(api *gin.RouterGroup ,userHandler *handler.UserHandler) {
	{
		api.POST("login", userHandler.Login)
		reqAuth := api.Group("", middleware.JwtAuthHandler())
		{
			userG := reqAuth.Group("user")
			{
				userG.GET("me", userHandler.Me)
				userG.PATCH("", userHandler.UpdateProfile)
			}
		}
	}
}
