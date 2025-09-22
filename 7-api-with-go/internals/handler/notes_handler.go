package handler

import (
	"errors"
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
	notes, err := h.s.GetNotes(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseList[[]dtosV1.NoteResponse]{
			Success: false,
			Message: "Failed To Fetch Notes!",
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.ApiResponseList[[]dtosV1.NoteResponse]{
		Success: true,
		Data:    notes,
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
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
			Success: false,
			Message: "Invalid Note Id",
			Data:    nil,
		})
		return
	}

	note, err := h.s.GetNote(ctx.Request.Context(), id)

	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			ctx.JSON(http.StatusNotFound, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
				Success: false,
				Message: "Note Not Found!",
				Data:    note,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
				Success: false,
				Message: "Failed To Fetch Note!",
				Data:    note,
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
		Success: true,
		Data:    note,
		Message: "Get Note Successfully!",
	})
}
