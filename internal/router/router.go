package router

import (
	"io"
	"net/http"

	"github.com/binhbeng/goex/config"
	// "github.com/binhbeng/goex/wire"
	// "github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/middleware"

	// "github.com/binhbeng/goex/internal/model"
	// "github.com/binhbeng/goex/internal/service"
	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
    UserHandler  *handler.UserHandler
}

func SetRouters(deps *RouterDeps) *gin.Engine {
	var engine *gin.Engine

	if config.Cfg.App.AppEnv == "production" {
		engine = ReleaseRouter()
		engine.Use(
		// 	middleware.RequestCostHandler(),
		// middleware.CustomRecovery(),
		)
	} else {
		engine = gin.New()
		engine.Use(
			gin.Logger(),
			// gin.Recovery(),
			// middleware.CustomLogger(),
			middleware.CustomRecovery(),
			// middleware.CorsHandler(),
		)
	}

	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	SetUserApiRoute(engine, deps.UserHandler)

	engine.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "404 not found",
		})
	})

	return engine
}

func ReleaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	engine := gin.New()

	return engine
}
