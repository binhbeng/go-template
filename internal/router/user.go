package router

import (
	"github.com/binhbeng/goex/internal/app"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserApiRoute(api *gin.RouterGroup) {
	userHandler:= app.NewUserModule().Handler()
	
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
