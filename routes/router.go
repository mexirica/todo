package routes

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	router.Run("0.0.0.0:" + port)
}
