package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internals/service"
)

type NotesHandler struct {
	s *service.NotesService
}

func NewNotesHandler(s *service.NotesService) *NotesHandler {
	return &NotesHandler{
		s,
	}
}

func (h NotesHandler) GetNotesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "revision Done"})
}
