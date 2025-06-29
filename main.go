// @title Iradat User Service API
// @version 1.0
// @description This is the User Service API for Iradat project.
// @host localhost:8000
// @BasePath /

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"iradat/profile/configs"
	"iradat/profile/internal/user"
	"log"
	"os"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Could not load .env file")
	}

	// Load necessary configurations
	configs.ReloadDatabaseConfig()
	configs.ReloadRedis()

	// Get HTTP port from environment or fallback
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	// Initialize Gin router
	router := gin.Default()

	// Register route groups
	user.HttpHandler(router)

	// Swagger docs endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Printf("✅ Starting HTTP server on port %s\n", httpPort)
	if err := router.Run(":" + httpPort); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
