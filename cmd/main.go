package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sonochiwa/news/configs"
	"github.com/sonochiwa/news/internal/global"
	"github.com/sonochiwa/news/internal/handlers"
	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/repositories"
	"github.com/sonochiwa/news/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.GlobalConfig

	gin.SetMode(getGinMode(config.Mode))

	gCtx := global.New(&config)

	{
		pg, err := postgres.New(config.Postgres)
		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}
		gCtx.Inst().Postgres = pg
	}

	repository := repositories.New(*gCtx.Inst())
	service := services.New(repository)
	handler := handlers.New(service).InitRoutes()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler: handler,
	}

	fmt.Printf("Server running on http://%s\n", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
		return
	}
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
