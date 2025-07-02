package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	APIGateway "auth/api_gateway"
	"auth/models"

	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Printf("Login attempt: %s | %s\n", email, password)

	// Placeholder logic
	if email == "admin@example.com" && password == "1234" {
		c.String(http.StatusOK, "Login successful!")
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials.")
	}
}

/* func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirm := c.PostForm("confirm_password")

	if password != confirm {
		c.String(http.StatusBadRequest, "Passwords do not match.")
		return
	}

	fmt.Printf("New user signup: %s\n", email)
	c.String(http.StatusOK, "Signup successful!")
} */

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirm := c.PostForm("confirm_password")

	if password != confirm {
		c.String(http.StatusBadRequest, "Passwords do not match.")
		return
	}

	user := models.User{
		Email:    email,
		Password: password, // NOTE: In real projects, always hash passwords!
	}

	// Convert user struct to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to encode user data.")
		return
	}

	fmt.Printf("New user signup JSON: %s\n", userJSON)

	APIGateway.API_Gateway_Json_Data_Request(userJSON)

}
