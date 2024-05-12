package handlers

import (
	"news/internal/services"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	service services.Services
}

func New(service services.Services) *gin.Engine {
	handler := handlers{service: service}

	router := gin.New()

	users := router.Group("/users")
	{
		users.GET("/", handler.getAllUsers)
		users.GET("/:id", handler.getUserByID)
	}

	return router
}
