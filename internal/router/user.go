package router

import (
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserApiRoute(e *gin.Engine ,userHandler *handler.UserHandler) {

	api := e.Group("api")
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
