// handlers/auth_handler.go
package handlers

import (
	"net/http"

	"api/user-api-v1/models"
	"api/user-api-v1/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var credentials models.LoginCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusBadRequest,
			Message:    "Login failed",
			Result:     gin.H{"error": "Invalid request payload", "result": nil},
		})
		return
	}

	if ah.authService.Authenticate(credentials.Email, credentials.Password) {
		c.JSON(http.StatusOK, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusOK,
			Message:    "Login successful",
			Result:     gin.H{"error": nil, "result": "Welcome " + credentials.Email + "!"},
		})
	} else {
		c.JSON(http.StatusUnauthorized, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid credentials",
			Result:     gin.H{"error": "Invalid credentials", "result": nil},
		})
	}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var credentials models.RegisterCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusBadRequest,
			Message:    "Registration failed",
			Result:     gin.H{"error": "Invalid request payload", "result": nil},
		})
		return
	}

	insert := ah.authService.Registration(credentials)

	if insert == 1 {
		c.JSON(http.StatusOK, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusOK,
			Message:    "Registration successful",
			Result:     gin.H{"error": nil, "result": "Welcome " + credentials.FullName + "!"},
		})
	} else if insert == -1 {
		c.JSON(http.StatusConflict, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusInternalServerError,
			Message:    "An error occurred while registering the user",
			Result:     gin.H{"error": "General error", "result": nil},
		})
	} else {
		c.JSON(http.StatusConflict, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusConflict,
			Message:    "User already exists",
			Result:     gin.H{"error": "User already exists", "result": nil},
		})
	}
}