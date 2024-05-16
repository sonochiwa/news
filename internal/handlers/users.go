package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	}
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	c.JSON(http.StatusOK, result)
}
