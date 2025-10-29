package routers

import (
	auth "github.com/binhbeng/goex/internal/controller/auth"
	"github.com/binhbeng/goex/internal/controller/user"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetAuthApiRoute(e *gin.Engine) {
	// version 1
	api := e.Group("api")
	{

		loginC := auth.NewLoginController()
		api.POST("login", loginC.Login)

		reqAuth := api.Group("", middleware.JwtAuthHandler())
		{
			userG := reqAuth.Group("user")
			{
				r := user.NewUserController()
				userG.GET("me", r.Me)
				userG.PATCH("", r.Update)
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
