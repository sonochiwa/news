package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) getAllCategories(c *gin.Context) {
	result, err := h.service.Categories.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}
