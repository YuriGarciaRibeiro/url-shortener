package main

import (
	"log"
	"os"

	"github.com/YuriGarciaRibeiro/url-shortener/internal/app/model"
	"github.com/YuriGarciaRibeiro/url-shortener/internal/app/service"
	"github.com/YuriGarciaRibeiro/url-shortener/internal/server/handlers"
	"github.com/YuriGarciaRibeiro/url-shortener/internal/server/middleware"
	"github.com/YuriGarciaRibeiro/url-shortener/internal/storage/postgres"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Swagger docs
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	docs "github.com/YuriGarciaRibeiro/url-shortener/docs"
)

// @title URL Shortener API
// @version 1.0
// @description A simple and efficient URL shortener built with Go and Gin.
// @termsOfService https://github.com/YuriGarciaRibeiro/url-shortener

// @contact.name Yuri Garcia Ribeiro
// @contact.url https://yurigarcia.dev/
// @contact.url https://yurigarcia.dev/
// @contact.email yuri@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Logger setup
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// PostgreSQL connection
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// DB migration
	if err := db.AutoMigrate(&model.Url{}); err != nil {
		logger.Fatal("Failed to migrate database", zap.Error(err))
	}

	// Dependency injection
	urlRepo := postgres.NewURLRepository(db)
	urlService := service.NewShortenerService(urlRepo)

	// Gin setup
	r := gin.New()
	r.Use(middleware.Logger(logger))
	r.Use(gin.Recovery())

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Register URL routes
	handlers.RegisterURLRoutes(r, urlService)

	// Swagger route
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
