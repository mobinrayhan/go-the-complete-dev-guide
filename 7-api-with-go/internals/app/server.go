package app

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "mobin.dev/internals/api/v1"
	"mobin.dev/internals/handlers"
	"mobin.dev/internals/repository"
	"mobin.dev/internals/service"
	"mobin.dev/pkg/config"
)

func RunServer(conf *config.Config, db *sql.DB) {
	fmt.Println(conf.Port)
	router := gin.Default()

	notesRepo := repository.NewNotesRepo(db)
	notesService := service.NewNotesService(notesRepo)
	notesHandler := handlers.NewNotesHandler(notesService)

	apiV1 := router.Group("/v1")
	{
		v1.RegisterNotesRoutes(apiV1, notesHandler)
	}

	router.Run(":" + strconv.Itoa(conf.Port))
}
