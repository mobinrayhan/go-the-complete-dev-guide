package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"mobin.dev/internals/model"
)

type NotesRepo struct{ db *sql.DB }

func NewNotesRepo(db *sql.DB) *NotesRepo {
	return &NotesRepo{
		db,
	}
}

func (r *NotesRepo) FetchAllNotes() ([]model.Note, error) {
	query := "SELECT * FROM notes"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query notes : %w", err)
	}

	defer rows.Close()

	var notes []model.Note

	for rows.Next() {
		var n model.Note
		var tagsByte []byte

		if err := rows.Scan(&n.Id, &n.UserId, &n.Title, &n.Body, &tagsByte, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan note : %w ", err)
		}

		if len(tagsByte) > 0 {
			if err := json.Unmarshal(tagsByte, &n.Tags); err != nil {
				return nil, fmt.Errorf("failed to unmarshal tags : %w ", err)
			}
		}
		notes = append(notes, n)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err : %w", err)
	}

	return notes, nil
}
