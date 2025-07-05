package controllers

import (
	"encoding/json"
	"fmt"
	"io"
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

/* func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Printf("Login attempt: %s | %s\n", email, password)

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

	fmt.Printf("Login Details: %s\n", userJSON)

	APIGateway.API_Gateway_Json_Data_Request_Login(userJSON)
} */

func Login(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userJSON, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode user data"})
		return
	}

	// Send the POST request
	resp, err := APIGateway.API_Gateway_Json_Data_Request_Login(userJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API Gateway request failed"})
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	// Unmarshal response into a map
	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON in response"})
		return
	}

	c.JSON(resp.StatusCode, responseMap)
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

/*
	 func Signup(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		confirm := c.PostForm("confirm_password")

		fmt.Printf("Raw Signup Form: email=%s, password=%s, confirm=%s\n", email, password, confirm)

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

		APIGateway.API_Gateway_Json_Data_Request_Signup(userJSON)

}
*/
func Signup(c *gin.Context) {
	var req struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Printf("Raw Signup Form: email=%s, password=%s, confirm=%s\n", req.Email, req.Password, req.ConfirmPassword)

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	// Convert to User struct
	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode user data"})
		return
	}

	responseJSON, statusCode, err := APIGateway.API_Gateway_Json_Data_Request_Signup(userJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API Gateway error"})
		return
	}

	// Define a struct matching SignupResponse
	type SignupResponse struct {
		Message      string `json:"message"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Role         string `json:"role"`
		Error        string `json:"error"`
	}

	var response SignupResponse
	if err := json.Unmarshal(responseJSON, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON response from API Gateway"})
		return
	}

	// Prepare response map
	responseMap := map[string]interface{}{
		"message":       response.Message,
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
		"role":          response.Role,
		"error":         response.Error,
	}

	c.JSON(statusCode, responseMap)
}

func ShowDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
