package main

import (
	"users-golang/controllers"
	"users-golang/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/user", controllers.UserCreation)
	r.GET("/users", controllers.UsersGetAll)
	r.GET("/user/:id", controllers.UserGetOne)
	r.PUT("/user/:id", controllers.UserUpdate)
	r.DELETE("/user/:id", controllers.UserDelete)

	r.Run()
}
