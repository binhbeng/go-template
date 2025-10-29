package routers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/wannanbigpig/gin-layout/config"
	// "github.com/wannanbigpig/gin-layout/internal/middleware"
	// "github.com/wannanbigpig/gin-layout/internal/pkg/errors"
)

func SetRouters() *gin.Engine {
	var engine *gin.Engine

	if false {
		engine = ReleaseRouter()
		// engine.Use(
		// 	middleware.RequestCostHandler(),
		// 	middleware.CustomLogger(),
		// 	middleware.CustomRecovery(),
		// 	middleware.CorsHandler(),
		// )
	} else {
		engine = gin.New()
		engine.Use(
			// middleware.RequestCostHandler(),
			gin.Logger(),
			// middleware.CustomRecovery(),
			// middleware.CorsHandler(),
		)
	}
	// set up trusted agents
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	// ping
	engine.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	SetAuthApiRoute(engine)

	engine.NoRoute(func(c *gin.Context) {
		fmt.Println("🔍 ~ SetRouters ~ internal/routers/router.go:48 ~ response2:")
	})

	return engine
}

func ReleaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	engine := gin.New()

	return engine
}
