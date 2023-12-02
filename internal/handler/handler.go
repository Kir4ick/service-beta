package handler

import (
	"beta/internal/services"
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
	ctx      *context.Context
}

func NewHandler(service *services.Service, ctx *context.Context) *Handler {
	return &Handler{services: service, ctx: ctx}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/voting/", h.voting)

	return router
}
