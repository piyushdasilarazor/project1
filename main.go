package main

import (
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(r)

	// Start the server
	r.Run(":8080")
}
