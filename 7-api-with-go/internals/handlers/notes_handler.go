package handlers

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internals/dtos"
	v1 "mobin.dev/internals/dtos/v1"
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

func (h *NotesHandler) GetAllNotesHandler(c *gin.Context) {
	notes, err := h.s.GetAllNotes()

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse[[]v1.NoteResponse]{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	resNotes := make([]v1.NoteResponse, len(notes))

	for i, note := range notes {
		resNotes[i] = v1.NoteResponse{
			Id:     note.Id,
			UserId: note.UserId, Title: note.Title,
			Tags:      note.Tags,
			Body:      note.Body,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt}
	}

	c.JSON(http.StatusOK, dtos.APIResponse[[]v1.NoteResponse]{
		Success: true,
		Data:    resNotes,
		Message: "Get All Notes Successfully!",
		Meta: dtos.ResponseMeta{
			Page:       1,
			Limit:      10,
			Total:      len(notes),
			TotalPages: int(math.Ceil(float64(len(notes)) / 10.0)),
		},
	})
}
