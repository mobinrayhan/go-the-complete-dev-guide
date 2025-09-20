package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mobin.dev/internals/dtos"
	dtosV1 "mobin.dev/internals/dtos/v1"
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

func (h *NotesHandler) GetNotesHandler(ctx *gin.Context) {
	notes, err := h.s.GetNotes()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponse[[]dtosV1.NoteResponse]{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.ApiResponse[[]dtosV1.NoteResponse]{
		Success: true,
		Data:    &notes,
		Message: "Get All Notes Successfully!",
		Meta: &dtos.ResponseMeta{
			Total:    len(notes),
			Page:     10,
			PerPages: 5,
			Limit:    20,
		},
	})
}

func (h *NotesHandler) GetNoteHandler(ctx *gin.Context) {
	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponse[dtosV1.NoteResponse]{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	note, err := h.s.GetNote(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponse[dtosV1.NoteResponse]{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.ApiResponse[dtosV1.NoteResponse]{
		Success: true,
		Data:    note,
		Message: "Get Note Successfully!",
	})
}
