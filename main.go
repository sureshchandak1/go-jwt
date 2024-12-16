package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sureshchandak1/go-jwt/controllers"
	"github.com/sureshchandak1/go-jwt/initializers"
	"github.com/sureshchandak1/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)

	r.Run()
}
