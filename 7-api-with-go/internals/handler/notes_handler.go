package handler

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mobin.dev/internals/dtos"
	dtosV1 "mobin.dev/internals/dtos/v1"
	"mobin.dev/internals/service"
	"mobin.dev/pkg/constants"
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
	page, errPage := strconv.Atoi(ctx.DefaultQuery("page", strconv.Itoa(constants.DefaultPage)))

	if errPage != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseList[[]dtosV1.NoteResponse]{
			Success: false,
			Message: "Invalid 'Page' Query Params!",
		})
		return
	}
	perPage, errPerPage := strconv.Atoi(ctx.DefaultQuery("perPage", strconv.Itoa(constants.DefaultPerPage)))

	if errPerPage != nil || perPage <= 0 {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseList[[]dtosV1.NoteResponse]{
			Success: false,
			Message: "Invalid 'Per Page' Query Params!",
		})
		return
	}
	skip := (page - 1) * perPage

	if perPage > constants.MaxPerPage {
		perPage = constants.MaxPerPage
	}

	notes, total, err := h.s.GetNotes(ctx.Request.Context(), skip, perPage)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseList[[]dtosV1.NoteResponse]{
			Success: false,
			Message: "Failed To Fetch Notes!",
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	hasNext := totalPages > page
	hasPrev := page > 1
	metaData := dtos.ResponseMeta{
		Total:      total,
		PerPage:    perPage,
		Page:       page,
		TotalPages: totalPages,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
	ctx.JSON(http.StatusOK, dtos.ApiResponseList[[]dtosV1.NoteResponse]{
		Success: true,
		Data:    notes,
		Message: "Get Notes Successfully!",
		Meta:    &metaData,
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
func (h *NotesHandler) CreateDummyNotesHandler(ctx *gin.Context) {
	var dummyNotesBody dtosV1.DummyNoteRequest

	if err := ctx.ShouldBindBodyWithJSON(&dummyNotesBody); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseSingle[any]{
			Message: "Size must be greater than 0",
			Success: false,
		})
		return
	}

	if err := h.s.CreateDummyNotes(ctx.Request.Context(), dummyNotesBody.Size); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseSingle[any]{
			Message: "Failed to create dummy notes",
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.ApiResponseSingle[dtosV1.DummyNoteRequest]{
		Message: "Dummy Notes Created Successfully",
		Data: &dtosV1.DummyNoteRequest{
			Size: dummyNotesBody.Size,
		},
		Success: true,
	})
}
