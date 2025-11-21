package routers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/binhbeng/goex/config"
	// "github.com/binhbeng/goex/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	var engine *gin.Engine

	if !config.C.App.Debug {
		engine = gin.New()
		engine.Use(
			// middleware.RequestCostHandler(),
			gin.Logger(),
			gin.Recovery(),
			// middleware.CustomLogger(),
			// middleware.CustomRecovery(),
			// middleware.CorsHandler(),
		)
	} else {
		engine = ReleaseRouter()
		engine.Use(
			// 	middleware.RequestCostHandler(),
			gin.Recovery(),
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

	SetAuthApiRoute(engine)

	engine.NoRoute(func(c *gin.Context) {
		fmt.Println("NoRoute")
	})

	return engine
}

func ReleaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	engine := gin.New()

	return engine
}
