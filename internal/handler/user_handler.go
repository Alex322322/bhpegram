package handler

import (
	"github.com/Alex322322/bhpegram/internal/domain"
	"github.com/Alex322322/bhpegram/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var u domain.CreateUserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.service.Create(c.Request.Context(), &u)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	var resp domain.UserResponse
	c.JSON(201, gin.H{"message": "User created", "user": resp})
}
