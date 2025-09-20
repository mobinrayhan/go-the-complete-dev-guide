package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"mobin.dev/internals/model"
)

type NotesRepo struct {
	db *sql.DB
}

func NewNotesRepo(db *sql.DB) *NotesRepo {
	return &NotesRepo{
		db,
	}
}

func (r *NotesRepo) FetchNotes() ([]model.Note, error) {
	rows, err := r.db.Query(`
				SELECT id, user_id, title, body, tags, created_at, updated_at
					FROM notes
				ORDER BY created_at DESC
		`)

	if err != nil {
		return nil, fmt.Errorf("failed to query db data %w ", err)
	}

	defer rows.Close()

	notes := make([]model.Note, 0)

	for rows.Next() {
		var note model.Note
		var tags json.RawMessage

		if err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &tags, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan rows : %w", err)
		}

		err := json.Unmarshal(tags, &note.Tags)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON : %w", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return notes, nil
}

func (r *NotesRepo) FetchNote(id int) (*model.Note, error) {
	fmt.Println(id)
	row := r.db.QueryRow(`SELECT id, user_id, title, body, tags, created_at, updated_at
					FROM notes WHERE id = $1`, id)

	var note model.Note
	var tags json.RawMessage

	if err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &tags, &note.CreatedAt, &note.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan note : %w ", err)
	}

	err := json.Unmarshal(tags, &note.Tags)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON : %w", err)
	}

	return &note, nil
}
