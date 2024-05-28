package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) getAllPosts(c *gin.Context) {
	filter := ""
	filter = c.Query("filter")
	result, err := h.service.Posts.GetAllPosts(filter)
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
