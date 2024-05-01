package news

import (
	"net/http"
	"news/internal/news"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := news.Setup()

	r.Use(gin.Logger())

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.ListenAndServe()
}
