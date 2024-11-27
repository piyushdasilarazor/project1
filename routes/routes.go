package routes

import (
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes maps the routes to service functions
func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/users", services.CreateUser)       // Create a new user
	router.GET("/users", services.GetUsers)          // Get all users
	router.GET("/users/:id", services.GetUser)       // Get a user by ID
	router.PUT("/users/:id", services.UpdateUser)    // Update a user by ID
	router.DELETE("/users/:id", services.DeleteUser) // Delete a user by ID
}
