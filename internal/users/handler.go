package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaocansi/simple-api/internal/helpers/errors"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *UserService
}

func Setup(s *gin.RouterGroup, db *gorm.DB) {
	service := NewUserService(db)
	handler := &UserHandler{service}

	r := s.Group("/users")

	r.POST("/", handler.createUser)
	r.POST("/sign-in", handler.signIn)
}

func (h *UserHandler) createUser(c *gin.Context) {
	type Payload struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		AvatarUrl string `json:"avatarUrl" binding:"required"`
	}

	var body Payload
	if err := c.ShouldBindJSON(&body); err != nil {
		errors.HttpError(c, err)
		return
	}

	user, err := h.service.createUser(CreateUser(body))

	if err != nil {
		errors.HttpError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"avatarUrl": user.AvatarUrl,
	})
}

func (h *UserHandler) signIn(c *gin.Context) {
	type Payload struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body Payload
	if err := c.ShouldBindJSON(&body); err != nil {
		errors.HttpError(c, err)
		return
	}

	signInResult, err := h.service.signIn(SignIn(body))
	if err != nil {
		errors.HttpError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": signInResult.AccessToken,
		"expiresIn":   signInResult.ExpiresIn,
	})
	c.SetCookie("accessToken", signInResult.AccessToken, 3600, "/", "", false, true)
}
