package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/utils"
)

func (h Handlers) updateUserPhoto(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := fmt.Sprintf("%s.%s", time.Now().Format("2006-01-02-04"), file.Filename)

	dst := "./public/" + filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

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

	err = h.service.Users.UpdateUserPhoto(int(result.ID), filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.String(http.StatusOK, "ok")

}

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
