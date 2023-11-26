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

func (h *Handler) voting(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	// Валидация
	input := request.Vote{}
	if err := ctx.ShouldBind(&input); err != nil {
		if errors.Is(err, io.EOF) {
			response.ErrorResponse(ctx, http.StatusBadRequest, "request body is empty")
		}

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			response.ValidationErrorResponse(ctx, verr)
		}

		return
	}
	go h.services.CreateVoting(&input)

	ctx.JSON(http.StatusOK, response.ResponseOk{Message: "ok"})
}
