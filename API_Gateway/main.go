package main

import (
	"apigateway/login"
	"apigateway/signup"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/signup", signup.SignupHandler)
	router.POST("/login", login.LoginHandler)

	// HTTPS only
	log.Println("🔐 Starting HTTPS server on https://172.20.78.91:443")
	log.Fatal(router.RunTLS(":443", "cert.pem", "key.pem"))
}
