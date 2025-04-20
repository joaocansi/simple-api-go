package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joaocansi/simple-api/internal/users"
)

func NewServerEngine(userHandler *users.UserHandler) *gin.Engine {
	r := gin.Default()
	s := r.Group("/api/v1")

	us := s.Group("/users")
	us.POST("/", userHandler.CreateUser)
	us.POST("/sign-in", userHandler.SignIn)

	return r
}
