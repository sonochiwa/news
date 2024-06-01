package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
