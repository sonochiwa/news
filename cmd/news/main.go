package news

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/internal/transport/rest"
)

func Run() {
	router := rest.Setup()

	router.Use(gin.Logger())

	server := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	server.ListenAndServe()
}
