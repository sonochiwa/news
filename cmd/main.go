package main

import (
	"fmt"
	"log"
	"net/http"
	"news/configs"
	"news/internal/transport/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.New()

	gin.SetMode(getGinMode(config.Mode))

	//gCtx := global.New(context.Background(), &config)

	router := rest.Setup()
	router.Use(gin.Logger())

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler: router,
	}

	fmt.Printf("Server running on %s://%s\n", config.Server.Protocol, server.Addr)

	server.ListenAndServe()
}

func getGinMode(key string) string {
	ginMode := map[string]string{
		"debug":   gin.DebugMode,
		"release": gin.ReleaseMode,
		"test":    gin.TestMode,
	}

	value, ok := ginMode[key]
	if !ok {
		log.Println("mode can be debug, release or test")
		log.Fatalf("unsupported mode %s", configs.New().Mode)
	}

	return value
}
