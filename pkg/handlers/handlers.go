package handlers

import (
	"github.com/Venukishore-R/CODECRAFT_BW_02/internal/app/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{Service: service}
}

type UserHandler interface {
	Create(ctx gin.Context)
	FindAll(ctx gin.Context)
	FindByEmail(ctx gin.Context)
	Update(ctx gin.Context)
	Delete(ctx gin.Context)
}
