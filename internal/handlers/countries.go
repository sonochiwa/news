package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) getAllCountries(c *gin.Context) {
	result, err := h.service.Countries.GetAllCountries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
