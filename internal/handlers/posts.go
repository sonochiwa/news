package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/utils"
)

func (h *Handlers) getAllPosts(c *gin.Context) {
	filter := c.Query("filter")
	category := c.Query("category")
	country := c.Query("country")

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

	result, err := h.service.Posts.GetAllPosts(&filter, &category, &country, &user.Language)
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

func (h *Handlers) deletePost(c *gin.Context) {
	parsedID, err := strconv.Atoi(c.Param("id"))

	err = h.service.Posts.DeletePost(parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")
}
