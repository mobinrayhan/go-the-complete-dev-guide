package service

import (
	"mobin.dev/internals/model"
	"mobin.dev/internals/repository"
)

type NotesService struct {
	r *repository.NotesRepo
}

func NewNotesService(r *repository.NotesRepo) *NotesService {
	return &NotesService{
		r,
	}
}

func (s NotesService) GetAllNotes() ([]model.Note, error) {
	notes, err := s.r.FetchAllNotes()

	if err != nil {
		return nil, err
	}

	return notes, nil
}
