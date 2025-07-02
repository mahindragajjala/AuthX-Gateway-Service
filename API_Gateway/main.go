package main

import (
	"apigateway/signup"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/signup", signup.SignupHandler)

	// HTTPS only
	log.Println("🔐 Starting HTTPS server on https://172.20.78.91:443")
	log.Fatal(router.RunTLS(":443", "cert.pem", "key.pem"))
}
