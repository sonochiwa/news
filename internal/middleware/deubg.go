package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func DebugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.Header)
		log.Println(c.RemoteIP())
	}
}
