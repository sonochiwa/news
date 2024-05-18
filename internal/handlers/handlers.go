package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/middleware"
	"github.com/sonochiwa/news/internal/services"
)

type Handlers struct {
	service services.Services
}

func New(service services.Services) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		api.GET("/posts", h.getAllPosts)
		api.GET("/categories", h.getAllCategories)
		api.GET("/languages", h.getAllLanguages)
	}

	authorizedApi := router.Group("/api") //middleware.AuthMiddleware()
	{
		authorizedApi.GET("/", h.getAllUsers)
		authorizedApi.GET("/:id", h.getUserByID)
	}

	return router
}
