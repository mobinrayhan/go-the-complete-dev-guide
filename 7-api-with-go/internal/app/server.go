package app

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internal/api"
	"mobin.dev/internal/repository"
	"mobin.dev/internal/service"
	"mobin.dev/pkg/config"
)

func RunServer(cfg config.Config) {
	router := gin.Default()

	// repos
	eventsRepo := repository.NewEventsRepo("Hello")

	// services
	eventsService := service.NewEventsService(*eventsRepo)

	api.RegisterEventsRouter(router, eventsService)

	router.Run(":" + cfg.Port)
}
