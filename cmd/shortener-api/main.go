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
	gormpostgres "gorm.io/driver/postgres" // Alias para evitar conflito de nomes
	"gorm.io/gorm"
)

func main() {
	// Configuração do Logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Conexão com o PostgreSQL (corrigido com alias)
	dsn := os.Getenv("DB_DSN") // Lê do ambiente
	db, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Migração do banco de dados
	if err := db.AutoMigrate(&model.Url{}); err != nil {
		logger.Fatal("Failed to migrate database", zap.Error(err))
	}
	

	// Inicializa as dependências
	urlRepo := postgres.NewURLRepository(db)
	urlService := service.NewShortenerService(urlRepo)

	// Configuração do Gin
	r := gin.New()

	// Middlewares
	r.Use(middleware.Logger(logger))
	r.Use(gin.Recovery())

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Registra rotas
	handlers.RegisterURLRoutes(r, urlService)

	// Inicia o servidor
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
