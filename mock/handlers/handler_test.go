
package handlers_test

import (
    "auth/mocks"
    "handlers"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAuth := mocks.NewMockAuthenticator(ctrl)
    mockAuth.EXPECT().Login("admin@example.com", "1234").Return(true)

    loginHandler := &handlers.LoginHandler{Auth: mockAuth}

    router := gin.Default()
    router.POST("/login", loginHandler.Login)

    body := "email=admin@example.com&password=1234"
    req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "Login successful!", w.Body.String())
}

func TestLogin_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAuth := mocks.NewMockAuthenticator(ctrl)
    mockAuth.EXPECT().Login("wrong@example.com", "wrongpass").Return(false)

    loginHandler := &handlers.LoginHandler{Auth: mockAuth}

    router := gin.Default()
    router.POST("/login", loginHandler.Login)

    body := "email=wrong@example.com&password=wrongpass"
    req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
    assert.Equal(t, "Invalid credentials.", w.Body.String())
}
