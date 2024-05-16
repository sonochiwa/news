package handlers

import (
	"github.com/sonochiwa/news/internal/middleware"
	"github.com/sonochiwa/news/internal/services"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	service services.Services
}

func New(service services.Services) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", middleware.AuthMiddleware())
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserByID)
		}
	}

	return router
}
