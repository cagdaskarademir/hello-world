package main

import (
	"log"
	"os"

	_ "github.com/cagdaskarademir/hello-world/docs"
	"github.com/cagdaskarademir/hello-world/internal/handlers"
	"github.com/cagdaskarademir/hello-world/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Hello World API
// @version         1.0
// @description     A simple Hello World API with health checks
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	router := gin.Default()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Health check endpoints
		v1.GET("/health/liveness", handlers.LivenessCheck)
		v1.GET("/health/readiness", handlers.ReadinessCheck)

		// Hello endpoint
		v1.GET("/say/:name", handlers.SayHello)
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
