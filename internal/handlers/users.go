package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/utils"
)

func (h Handlers) getAllUsers(c *gin.Context) {
	result, err := h.service.Users.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}

func (h Handlers) getUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	result, err := h.service.Users.GetUserByID(int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h Handlers) getMe(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
		c.Abort()
		return
	}

	tokenParts := strings.Split(header, " ")
	if len(tokenParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
		c.Abort()
		return
	}

	username, err := utils.ParseToken(tokenParts[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	result, err := h.service.Users.GetUserByLogin(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
