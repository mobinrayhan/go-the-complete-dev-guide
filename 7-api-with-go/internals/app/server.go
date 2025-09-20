package app

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "mobin.dev/internals/api/v1"
	"mobin.dev/internals/handler"
	"mobin.dev/internals/repository"
	"mobin.dev/internals/service"
	"mobin.dev/pkg/config"
)

func RunServer(c *config.Config, db *sql.DB) {

	r := gin.Default()

	groupV1 := r.Group("/v1")
	notesRepo := repository.NewNotesRepo(db)
	notesService := service.NewNotesService(notesRepo)
	notesHandler := handler.NewNotesHandler(notesService)

	{
		v1.RegisterNotesRoutes(groupV1, *notesHandler)
	}

	fmt.Printf("App Running One : http://localhost:%d\n", c.Port)
	r.Run(":" + strconv.Itoa(c.Port))

}
