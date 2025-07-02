package signup

import (
	"apigateway/grpc"
	"apigateway/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup handler with validation
func SignupHandler(c *gin.Context) {
	var req models.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	grpc.StartGRPCClient(req)

	log.Printf("Received signup: %s", req.Email)
	c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
}
