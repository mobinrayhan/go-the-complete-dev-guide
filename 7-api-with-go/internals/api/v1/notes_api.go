package apiV1

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internals/handler"
)

func RegisterNotesRoutes(routes *gin.RouterGroup, h handler.NotesHandler) {
	notes := routes.Group("/notes")

	{
		notes.GET("/", h.GetNotesHandler)
		notes.POST("/random-notes", h.CreateDummyNotesHandler)
		notes.GET("/:id", h.GetNoteHandler)
	}
}
