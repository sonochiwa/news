package handlers

import (
	"github.com/gin-gonic/contrib/static"
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

	router.MaxMultipartMemory = 8 << 20

	router.Use(static.Serve("/images", static.LocalFile("./public", true)))

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
		api.POST("/posts", h.newPost)
		api.GET("/categories", h.getAllCategories)
	}

	authorizedApi := router.Group("/api", middleware.AuthMiddleware())
	{
		authorizedApi.POST("/upload/img", h.updateUserPhoto)
		authorizedApi.GET("/users", h.getAllUsers)
		authorizedApi.GET("/users/:id", h.getUserByID)
		authorizedApi.GET("/users/me", h.getMe)
	}

	return router
}
