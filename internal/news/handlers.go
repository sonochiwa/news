package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("/var/www/html/web/templates/*")

	router.GET("/", indexHandler)
	router.GET("/sign-in", signInHandler)
	router.GET("/sign-up", signUpHandler)
	router.GET("/profile", profileHandler)
	router.GET("/my", myHandler)

	return router
}

func indexHandler(c *gin.Context) {
	data := gin.H{
		"title": "News | Главная",
	}

	c.HTML(http.StatusOK, "index.tmpl", data)
}

func signInHandler(c *gin.Context) {
	data := gin.H{
		"title": "News | Вход",
	}

	c.HTML(http.StatusOK, "sign-in.tmpl", data)
}

func signUpHandler(c *gin.Context) {
	data := gin.H{
		"title": "News | Регистрация",
	}

	c.HTML(http.StatusOK, "sign-up.tmpl", data)
}

func profileHandler(c *gin.Context) {

}

func myHandler(c *gin.Context) {

}
