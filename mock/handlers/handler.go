
package handlers

import (
    "auth"
    "github.com/gin-gonic/gin"
    "net/http"
)

type LoginHandler struct {
    Auth auth.Authenticator
}

func (lh *LoginHandler) Login(c *gin.Context) {
    email := c.PostForm("email")
    password := c.PostForm("password")

    if lh.Auth.Login(email, password) {
        c.String(http.StatusOK, "Login successful!")
    } else {
        c.String(http.StatusUnauthorized, "Invalid credentials.")
    }
}
