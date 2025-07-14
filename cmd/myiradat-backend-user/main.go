// @title Iradat User Service API
// @version 1.0
// @description This is the User Service API for Iradat project.
// @host localhost:8000
// @BasePath /

package main

import (
	"fmt"
	"log"
	"myiradat-backend-auth/docs"
	_ "myiradat-backend-auth/docs"
	configs2 "myiradat-backend-auth/internal/configs"
	"myiradat-backend-auth/internal/user"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ Could not load .env file")
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	envFile := fmt.Sprintf(".env_%s", env)
	if err := godotenv.Overload(envFile); err != nil {
		log.Printf("⚠️ Could not load %s\n", envFile)
	}

	// Load configurations
	configs2.ReloadDatabaseConfig()
	//configs2.ReloadRedis()

	// Get HTTP port
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	docs.SwaggerInfo.Host = "localhost:" + httpPort
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Iradat User Service API"
	docs.SwaggerInfo.Description = "This is the User Service API for Iradat project."
	docs.SwaggerInfo.Version = "1.0"

	router := gin.Default()
	user.HttpHandler(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("✅ Starting HTTP server on port %s (ENV=%s)\n", httpPort, env)
	if err := router.Run(":" + httpPort); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
