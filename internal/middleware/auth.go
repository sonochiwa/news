package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		login, err := utils.ParseToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("login", login)
	}
}
