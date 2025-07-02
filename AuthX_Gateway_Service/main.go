package main

import (
	"auth/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Static and HTML templates
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login", controllers.Login)
	r.GET("/signup", controllers.ShowSignupPage)
	r.POST("/signup", controllers.Signup)

	r.Run("172.20.78.91:8080")
}
