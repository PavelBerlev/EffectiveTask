package main

import (
	"EffectiveTask/internal/config"
	subHandler "EffectiveTask/internal/handler/subscribe"
	subRepo "EffectiveTask/internal/repository/subscribe"
	subService "EffectiveTask/internal/service/subscribe"
	"EffectiveTask/pkg/internalsql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "EffectiveTask/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title EffectiveTask API
// @version 1.0
// @description API для управления подписками
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// For Swagger
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectPostgreSQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger := log.New(os.Stdout, "effective: ", log.LstdFlags|log.Lmicroseconds)
	subscribeRepo := subRepo.NewRepository(db, logger)
	subService := subService.NewService(cfg, subscribeRepo, logger)
	subscribeHandler := subHandler.NewHandler(r, subService, logger)

	subscribeHandler.RouteList()

	server := fmt.Sprintf(":%s", cfg.AppPort)
	r.Run(server)
}
