package router

import (
	"io"
	"net/http"

	"github.com/binhbeng/goex/config"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type RouterDeps struct {
	UserHandler *handler.UserHandler
}

func SetRouters(deps *RouterDeps) *gin.Engine {
	var engine *gin.Engine

	if config.Cfg.App.AppEnv == "production" {
		engine = ReleaseRouter()
		engine.Use(
			gin.Logger(),
			gin.Recovery(),
		)
	} else {
		engine = gin.New()
		engine.Use(
			middleware.CustomLogger(config.Cfg.App.EnableBodyLog),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
		engine.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	api := engine.Group("/api")

	if err != nil {
		panic(err)
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	SetUserApiRoute(api, deps.UserHandler)

	return engine
}

func ReleaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	engine := gin.New()

	return engine
}
