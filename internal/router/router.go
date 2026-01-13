package router

import (
	"io"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/binhbeng/goex/config"
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/gin-gonic/gin"
	// "github.com/swaggo/files"
	// "github.com/swaggo/gin-swagger"
)

func SetRouters() *gin.Engine {
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
		engine.GET("/api/docs", func(c *gin.Context) {
			htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
				SpecURL:       "./docs/swagger.json",
				CustomOptions: scalar.CustomOptions{
					PageTitle: "GOEX API",
				},
				DarkMode:   false,
				IsEditable: false,
				WithDefaultFonts: true,
			})

			if err != nil {
				c.String(500, "failed to generate API reference: %v", err)
				return
			}

			c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
		})
	}

	if(config.Cfg.App.Socket) {
		engine.GET("/ws", func(c *gin.Context) {
			data.HandleWebSocket(c)
		})
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

	SetUserApiRoute(api)

	return engine
}

func ReleaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	engine := gin.New()

	return engine
}
