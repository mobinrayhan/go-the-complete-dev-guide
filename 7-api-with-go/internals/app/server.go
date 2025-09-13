package app

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"mobin.dev/pkg/config"
)

func RunServer(conf *config.Config) {
	fmt.Println(conf.Port)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":" + strconv.Itoa(conf.Port))
}
