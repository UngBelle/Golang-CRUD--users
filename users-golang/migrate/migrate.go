package main

import (
	"users-golang/initializers"
	"users-golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Users{})
}
