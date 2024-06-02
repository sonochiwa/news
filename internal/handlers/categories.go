package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/utils"
)

func (h *Handlers) getAllCategories(c *gin.Context) {
	user := &models.UserMe{}

	header := c.GetHeader("Authorization")
	if header == "" {
		user.Language = "ru"
	} else {
		tokenParts := strings.Split(header, " ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			c.Abort()
			return
		}

		login, err := utils.ParseToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		user, _ = h.service.Users.GetUserByLogin(login)
	}

	result, err := h.service.Categories.GetAllCategories(user.Language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
