package services

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// In-memory user store
var users = make(map[int]models.User)

// Global variable for assigning IDs
var nextID = 1

// CreateUser adds a new user
func CreateUser(c *gin.Context) {
	var newUser models.User

	// Validate incoming JSON
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign ID and increment for the next user
	newUser.ID = nextID
	nextID++

	// Store the user in the map
	users[newUser.ID] = newUser

	// Return the created user
	c.JSON(http.StatusCreated, newUser)
}

// GetUsers retrieves all users
func GetUsers(c *gin.Context) {
	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}
	c.JSON(http.StatusOK, userList)
}

// GetUser retrieves a user by ID
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if user, exists := users[id]; exists {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

// UpdateUser modifies an existing user by ID
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser models.User

	// Validate incoming JSON
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	if user, exists := users[id]; exists {
		user.Name = updatedUser.Name
		user.Age = updatedUser.Age
		users[id] = user

		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

// DeleteUser removes a user by ID
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if _, exists := users[id]; exists {
		delete(users, id)
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}
