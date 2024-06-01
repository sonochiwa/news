package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/models"
)

func (h *Handlers) getAllPosts(c *gin.Context) {
	var filter, category string
	filter = c.Query("filter")
	category = c.Query("category")

	result, err := h.service.Posts.GetAllPosts(&filter, &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result == nil {
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handlers) newPost(c *gin.Context) {
	var input models.NewPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Posts.NewPost(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")
}
