package main

import (
	"fmt"
	"log"
	"net/http"
	"news/configs"
	"news/internal/handlers"
	"news/internal/instances/postgres"
	"news/internal/repositories"
	"news/internal/services"
	"news/pkg/global"

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
	handler := handlers.New(service)

	handler.Use(gin.Logger())

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler: handler,
	}

	fmt.Printf("Server running on %s://%s\n", config.Server.Protocol, server.Addr)

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
