package handlers

import (
	"net/http"
	"news/pkg/middleware"
	"news/pkg/services"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	service services.Services
}

func New(service services.Services) *gin.Engine {
	handler := handlers{service: service}

	router := gin.New()

	router.Use(gin.Logger())

	router.POST("/sign-in", handler.signIn)
	router.POST("/sign-up", handler.signUp)

	auth := router.Group("/", middleware.AuthMiddleware())
	{
		auth.GET("/protected", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(http.StatusOK, gin.H{"message": "protected route accessed by " + username})
		})
	}

	users := auth.Group("/users")
	{
		users.GET("/", handler.getAllUsers)
		users.GET("/:id", handler.getUserByID)
	}

	return router
}
