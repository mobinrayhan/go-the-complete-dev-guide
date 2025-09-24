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
	ErrNoteNotFound       = errors.New("note not found")
	ErrNotesFetchFailed   = errors.New("failed to fetch notes")
	ErrorDummyNotesInsert = errors.New("failed to insert dummy notes")
)

type NotesService struct {
	r *repository.NotesRepo
}

func NewNotesService(r *repository.NotesRepo) *NotesService {
	return &NotesService{
		r,
	}
}

func (s *NotesService) GetNotes(ctx context.Context, skip, perPage int) ([]dtosV1.NoteResponse, int, error) {
	notes, total, err := s.r.FetchNotes(ctx, skip, perPage)

	if err != nil {
		return nil, total, fmt.Errorf("%w : %v", ErrNotesFetchFailed, err)
	}

	return dtosV1.ToNoteResponses(notes), total, nil
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
func (s *NotesService) CreateDummyNotes(ctx context.Context, size int) error {

	err := s.r.PostDummyNotes(ctx, size)

	if err != nil {
		return fmt.Errorf("create dummy notes failed: %w", err)
	}

	return nil
}
