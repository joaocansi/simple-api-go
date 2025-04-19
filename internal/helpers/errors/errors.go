package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceError struct {
	Code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewError(code int, status, message string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Status:  status,
		Message: message,
	}
}

func UserNotFound() *ServiceError {
	return NewError(http.StatusNotFound, "USER_NOT_FOUND", "usuário não encontrado")
}

func UserAlreadyExists() *ServiceError {
	return NewError(http.StatusConflict, "USER_ALREADY_EXISTS", "usuário já cadastrado")
}

func WrongUserCredentials() *ServiceError {
	return NewError(http.StatusUnauthorized, "WRONG_USER_CREDENTIALS", "email/senha não está correto")
}

func InternalError() *ServiceError {
	return NewError(http.StatusInternalServerError, "INTERNAL_ERROR", "erro interno")
}

func HttpError(c *gin.Context, err error) {
	if serviceErr, ok := err.(*ServiceError); ok {
		c.JSON(serviceErr.Code, err)
		c.Abort()
		return
	}
	HttpError(c, InternalError())
}
