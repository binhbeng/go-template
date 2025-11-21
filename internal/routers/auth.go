package routers

import (
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/service"
	"github.com/gin-gonic/gin"
)

func SetAuthApiRoute(e *gin.Engine) {

	// userService := service.NewUserService(service.NewService(), &model.User{})
	userService := service.NewUserService(&model.User{})
	userHandler := handler.NewUserHandler(handler.NewHandler(), userService)

	api := e.Group("api")
	{

		api.POST("login", userHandler.Login)

		reqAuth := api.Group("", middleware.JwtAuthHandler())
		{
			userG := reqAuth.Group("user")
			{
				userG.GET("me", userHandler.Me)
				userG.PATCH("", userHandler.Update)
			}

			// permissions := reqAuth.Group("permission")
			// {
			// 	r := admin_v1.NewPermissionController()
			// 	permissions.POST("edit", r.Edit)
			// 	permissions.GET("list", r.List)
			// }

		}
	}
}
