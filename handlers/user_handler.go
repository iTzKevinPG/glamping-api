package handlers

import (
	"api/user-api-v1/models"
	"api/user-api-v1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (ah *UserHandler) GetUser(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusBadRequest,
			Message:    "Get user failed",
			Result:     gin.H{"error": "Invalid request payload", "result": nil},
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusBadRequest,
			Message:    "Get user failed",
			Result:     gin.H{"error": "Invalid user ID", "result": nil},
		})
		return
	}

	user := ah.userService.GetUserById(idInt)

	if user == nil {
		c.JSON(http.StatusInternalServerError, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusInternalServerError,
			Message:    "Get user failed",
			Result:     gin.H{"error": "An error ocurred when consult the user", "result": nil},
		})
		return
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		APIVersion: "1.0",
		StatusCode: http.StatusOK,
		Message:    "Get user successful",
		Result:     gin.H{"error": nil, "result": gin.H{"user": user}},
	})
}

func (ah *UserHandler) UpdateUser(c *gin.Context) {
	var credentials models.UpdateUser

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusBadRequest,
			Message:    "Update user failed",
			Result:     gin.H{"error": "Invalid request payload", "result": nil},
		})
		return
	}

	if credentials.PayMethodId == 0 {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusBadRequest,
			Message:    "Update user failed",
			Result:     gin.H{"error": "Invalid request payload: payMethodId", "result": nil},
		})
		return
	}

	updated := ah.userService.UpdateUserById(credentials)

	if updated == 1 {
		c.JSON(http.StatusOK, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusOK,
			Message:    "User Updated successful",
			Result:     gin.H{"error": nil, "result": gin.H{"user": credentials}},
		})
	} else if updated == -1 {
		c.JSON(http.StatusConflict, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusInternalServerError,
			Message:    "An error occurred while update the user",
			Result:     gin.H{"error": "General error", "result": nil},
		})
	} else {
		c.JSON(http.StatusNotFound, models.GenericResponse{
			APIVersion: "1.0",
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
			Result:     gin.H{"error": "User not found", "result": nil},
		})
	}
}
