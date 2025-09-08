package app

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internal/api"
	"mobin.dev/pkg/config"
)

func RunServer(cfg config.Config) {
	router := gin.Default()
	api.RegisterRouter(router)
	router.Run(":" + cfg.Port)
}
