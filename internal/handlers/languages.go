package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) getAllLanguages(c *gin.Context) {
	result, err := h.service.Languages.GetAllLanguages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
