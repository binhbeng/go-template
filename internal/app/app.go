package app

import (
	"io"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/binhbeng/goex/config"
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/middleware"
	"github.com/binhbeng/goex/internal/router"
	"github.com/binhbeng/goex/pkg/kafka"
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine       *gin.Engine
	OrderHandler *handler.OrderHandler
	UserHandler  *handler.UserHandler
}

func NewApp() *App {
	config.Load()
	data.InitData()

	kafkaProducer := kafka.NewKafkaProducer([]string{"localhost:9092"})

	orderModule := NewOrderModule(kafkaProducer)
	userModule := NewUserModule()

	app := &App{
		OrderHandler: orderModule.Handler(),
		UserHandler:  userModule.Handler(),
	}

	app.setupRouter()

	return app
}

func (a *App) setupRouter() {
	var engine *gin.Engine

	if config.Cfg.App.AppEnv == "production" {
		engine = a.releaseRouter()
		engine.Use(
			gin.Logger(),
			gin.Recovery(),
		)
	} else {
		engine = gin.New()
		engine.Use(
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)

		engine.GET("/api/docs", func(c *gin.Context) {
			htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
				SpecURL: "./docs/swagger.json",
				CustomOptions: scalar.CustomOptions{
					PageTitle: "GOEX API",
				},
				DarkMode:         false,
				IsEditable:       false,
				WithDefaultFonts: true,
			})

			if err != nil {
				c.String(500, "failed to generate API reference: %v", err)
				return
			}

			c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
		})
	}

	if config.Cfg.App.Socket {
		engine.GET("/ws", func(c *gin.Context) {
			data.HandleWebSocket(c)
		})
	}

	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	api := engine.Group("/api")
	router.SetOrderApiRoute(api, a.OrderHandler)
	router.SetUserApiRoute(api, a.UserHandler)

	a.Engine = engine
}

func (a *App) releaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	return gin.New()
}
