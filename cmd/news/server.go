package news

import (
	"net/http"
	"news/internal/news"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := news.Setup()

	router.Use(gin.Logger())

	server := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	server.ListenAndServe()
}
