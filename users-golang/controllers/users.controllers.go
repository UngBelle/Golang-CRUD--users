package controllers

import (
	"users-golang/initializers"
	"users-golang/models"

	"github.com/gin-gonic/gin"
)

func UserCreation(c *gin.Context) {
	// Get data
	var users struct {
		Name       string
		Email      string
		Age        string
		DayofBirth string
	}
	c.Bind(&users)

	// Create a user
	user := models.Users{
		Name:       users.Name,
		Email:      users.Email,
		Age:        users.Age,
		DayofBirth: users.DayofBirth,
	}
	// Store a user
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return
	c.JSON(200, gin.H{
		"users": user,
	})
}

func UsersGetAll(c *gin.Context) {
	// Get users
	var users []models.Users
	initializers.DB.Find(&users)

	// Return
	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserGetOne(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Get a user by ID
	var user models.Users
	initializers.DB.Find(&user, id)

	// Return
	c.JSON(200, gin.H{
		"users": user,
	})

}

func UserUpdate(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Get a user
	var users struct {
		Name       string
		Email      string
		Age        string
		DayofBirth string
	}
	c.Bind(&users)

	// Find the user
	var user models.Users
	initializers.DB.Find(&user, id)

	// Update the user
	initializers.DB.Model(&user).Updates(models.Users{
		Name:       users.Name,
		Email:      users.Email,
		Age:        users.Age,
		DayofBirth: users.DayofBirth,
	})

	// Return
	c.JSON(200, gin.H{
		"users": user,
	})
}

func UserDelete(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Delete a user
	initializers.DB.Delete(&models.Users{}, id)

	// Return
	c.Status(200)
}
