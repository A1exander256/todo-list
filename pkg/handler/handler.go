package handler

import (
	"github.com/a1exander256/todo/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
	log     *logrus.Logger
}

func NewHandler(service *service.Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) InitRoutes(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	router := gin.New()

	todo := router.Group("/todo")
	{
		todo.POST("/", h.createItem)
		todo.GET("/:id", h.getItemById)
	}
	return router
}
