package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fauzannursalma/mineport/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Connect Database
	config.ConnectDB(cfg)

	// Migrate Database
	config.MigrateDB()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := ":8080"
	fmt.Println("Server running on port", port)
	fmt.Println("Link to server: http://localhost" + port)

  // check jwt token
	fmt.Println("Checking JWT token..." + cfg.JWTSecret)

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
