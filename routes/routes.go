package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./web/templates/*")

	r.GET("/", func(c *gin.Context) {
		data := gin.H{
			"title": "News | Главная",
		}

		c.HTML(http.StatusOK, "index.tmpl", data)
	})

	r.GET("/sign-up", func(c *gin.Context) {
		data := gin.H{
			"title": "News | Регистрация",
		}

		c.HTML(http.StatusOK, "sign-up.tmpl", data)
	})

	r.GET("/sign-in", func(c *gin.Context) {
		data := gin.H{
			"title": "News | Вход",
		}

		c.HTML(http.StatusOK, "sign-in.tmpl", data)
	})

	return r
}
