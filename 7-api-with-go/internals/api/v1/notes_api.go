package v1

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internals/handlers"
)

func RegisterNotesRoutes(router *gin.RouterGroup, h *handlers.NotesHandler) {
	notes := router.Group("/notes")
	{
		notes.GET("/", h.GetAllNotesHandler)
	}
}
