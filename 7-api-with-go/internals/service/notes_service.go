package service

import (
	dtosV1 "mobin.dev/internals/dtos/v1"
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

func (s *NotesService) GetNotes() ([]dtosV1.NoteResponse, error) {

	notes, err := s.r.FetchNotes()
	notesRes := make([]dtosV1.NoteResponse, len(notes))

	if err != nil {
		return nil, err
	}

	for index, note := range notes {
		notesRes[index] = dtosV1.NoteResponse{
			Id:        note.Id,
			UserId:    note.UserId,
			Title:     note.Title,
			Body:      note.Body,
			Tags:      note.Tags,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		}
	}

	return notesRes, nil
}

func (s *NotesService) GetNote(id int) (*dtosV1.NoteResponse, error) {

	note, err := s.r.FetchNote(id)

	if err != nil {
		return nil, err
	}

	noteRes := dtosV1.NoteResponse{
		Id:        note.Id,
		UserId:    note.UserId,
		Title:     note.Title,
		Body:      note.Body,
		Tags:      note.Tags,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}

	return &noteRes, nil
}
