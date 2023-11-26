package response

import (
	"beta/pkg/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}

func ValidationErrorResponse(c *gin.Context, errors validator.ValidationErrors) {
	c.AbortWithStatusJSON(http.StatusBadRequest, validation.GetErrors(errors))
}
