package handler

import (
	"beta/internal/request"
	"beta/internal/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

func (h *Handler) voting(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var input request.Vote
	if err := c.ShouldBind(&input); err != nil {
		if errors.Is(err, io.EOF) {
			response.ErrorResponse(c, http.StatusBadRequest, "request body is empty")
			return
		}

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			response.ValidationErrorResponse(c, verr)
			return
		}
	}

	c.JSON(http.StatusOK, response.ResponseOk{Message: "ok"})
}
