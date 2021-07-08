package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/advertising/delivery"
	"github.com/polyanimal/advertising/internal/advertising/repository"
	"github.com/polyanimal/advertising/internal/advertising/usecase"
	"github.com/polyanimal/advertising/internal/middleware"
	"log"
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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func NewServer() *App {
	connStr, connected := os.LookupEnv("DB_CONNECT")
	if !connected {
		fmt.Println(os.Getwd())
		log.Fatal("Failed to read DB connection data")
	}

	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	advertisingRepo := repository.NewAdvertisingRepo(dbpool)
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
	config.AllowOrigins = []string{"http://localhost" + ":" + port}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.Use(gin.Recovery())
	delivery.RegisterHTTPEndpoints(router, app.advertisingUC, app.validationMiddleware)

	app.server = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := app.server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to listen and serve: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.server.Shutdown(ctx)
}
