package signup

import (
	"apigateway/grpc"
	"apigateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignupHandler handles HTTP request and calls the gRPC SignupUser method
func SignupHandler(c *gin.Context) {
	var req models.SignupRequest

	// Validate JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call gRPC Signup service
	response, err := grpc.StartGRPCClient_Signup(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gRPC call failed"})
		return
	}
	// Convert protobuf response to JSON response
	c.JSON(http.StatusOK, gin.H{
		"message":       response.Message,
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
		"role":          response.Role,
		"error":         response.Error,
	})

}
