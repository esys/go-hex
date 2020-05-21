package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-hex/internal/user/domain"
)

type UserHandler interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type userHandler struct {
	userService domain.UserService
}

func NewUserHandler(service domain.UserService) UserHandler {
	return &userHandler{userService: service}
}

func (h *userHandler) Get(c *gin.Context) {
	if users, err := h.userService.FindAllUsers(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (h *userHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	if user, err := h.userService.FindUserById(id); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (h *userHandler) Create(c *gin.Context) {
	var json domain.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if err := h.userService.CreateUser(&json); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	} else {
		c.JSON(http.StatusCreated, CreatedResponse{json})
	}
}
