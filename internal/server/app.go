package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/advertising/internal/advertising"
	advertisingHttp "github.com/polyanimal/advertising/internal/advertising/delivery/http"
	localstorage "github.com/polyanimal/advertising/internal/advertising/repository/localstorage"
	"github.com/polyanimal/advertising/internal/advertising/usecase"
	middleware "github.com/polyanimal/advertising/internal/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	server               *http.Server
	advertisingUC        advertising.UseCase
	validationMiddleware middleware.Validation
}

func NewServer() *App {
	advertisingRepo := localstorage.NewAdvertisingRepo()
	advertisingUC := usecase.NewAdvertisingUC(advertisingRepo)

	validationMiddleware := middleware.NewValidationMiddleware()

	return &App{
		advertisingUC:        advertisingUC,
		validationMiddleware: validationMiddleware,
	}
}

func (app *App) Run(port string) error {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.Use(gin.Recovery())
	advertisingHttp.RegisterHTTPEndpoints(router, app.advertisingUC)

	app.server = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
8080
	return app.server.Shutdown(ctx)
}
