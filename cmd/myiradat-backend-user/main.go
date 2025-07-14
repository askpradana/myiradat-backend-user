// @title Iradat User Service API
// @version 1.0
// @description This is the User Service API for Iradat project.
// @host localhost:8000
// @BasePath /

package main

import (
	"log"
	// "myiradat-backend-auth/docs"
	config "myiradat-backend-auth/internal/configs"
	"myiradat-backend-auth/internal/user"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	// Load configurations
	config.ReloadDatabaseConfig()
	// config.ReloadRedis()

	// Get HTTP port
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	// docs.SetupSwagger(httpPort)

	router := gin.Default()
	user.HttpHandler(router)

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("✅ Starting HTTP server on port %s (ENV=%s)\n", httpPort, os.Getenv("ENV"))
	if err := router.Run(":" + httpPort); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
