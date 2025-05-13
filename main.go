package main

import (
	routes "github.com/akhil/golang-jwt-project/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		port = "10000"
	}

	if os.Getenv("SECRET_KEY") == "" {
		log.Fatal("SECRET_KEY is required")
	}

	if os.Getenv("MONGODB_URL") == "" {
		log.Fatal("MONGODB_URL is required")
	}

	router := gin.Default()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "JWT auth demo is running",
			"docs":    "POST /users/signup, POST /users/login, GET /users/me",
		})
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "golang-jwt-demo",
		})
	})

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	log.Fatal(router.Run(":" + port))
}
