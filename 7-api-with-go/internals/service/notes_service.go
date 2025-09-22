package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	dtosV1 "mobin.dev/internals/dtos/v1"
	"mobin.dev/internals/repository"
)

var (
	ErrNoteNotFound     = errors.New("note not found")
	ErrNotesFetchFailed = errors.New("failed to fetch notes")
)

type NotesService struct {
	r *repository.NotesRepo
}

func NewNotesService(r *repository.NotesRepo) *NotesService {
	return &NotesService{
		r,
	}
}

func (s *NotesService) GetNotes(ctx context.Context) ([]dtosV1.NoteResponse, error) {
	notes, err := s.r.FetchNotes(ctx)

	if err != nil {
		return nil, fmt.Errorf("%w : %v", ErrNotesFetchFailed, err)
	}

	return dtosV1.ToNoteResponses(notes), nil
}

func (s *NotesService) GetNote(ctx context.Context, id int) (*dtosV1.NoteResponse, error) {

	note, err := s.r.FetchNote(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}

	noteRes := dtosV1.ToNoteResponse(*note)

	return &noteRes, nil
}
