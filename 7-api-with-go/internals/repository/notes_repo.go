package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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

func (r *NotesRepo) FetchNotes(ctx context.Context) ([]model.Note, int, error) {
	var total int
	totalRows := r.db.QueryRowContext(ctx, `SELECT COUNT(*) AS total FROM notes`)

	if err := totalRows.Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to scan total rows : %w", err)
	}

	rows, err := r.db.QueryContext(ctx, `
				SELECT id, user_id, title, body, tags, created_at, updated_at
					FROM notes
				ORDER BY created_at DESC
		`)

	if err != nil {
		return nil, total, fmt.Errorf("failed to query db data %w ", err)
	}

	defer rows.Close()

	notes := make([]model.Note, 0)

	for rows.Next() {
		var note model.Note
		var tags json.RawMessage

		if err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &tags, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, total, fmt.Errorf("failed to scan rows : %w", err)
		}

		err := json.Unmarshal(tags, &note.Tags)
		if err != nil {
			return nil, total, fmt.Errorf("error unmarshalling JSON : %w", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, total, fmt.Errorf("row iteration error: %w", err)
	}

	return notes, total, nil
}

func (r *NotesRepo) FetchNote(ctx context.Context, id int) (*model.Note, error) {
	fmt.Println(id)
	row := r.db.QueryRowContext(ctx, `SELECT id, user_id, title, body, tags, created_at, updated_at
					FROM notes WHERE id = $1`, id)

	var note model.Note
	var tags json.RawMessage

	err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &tags, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("failed to scan note: %w", err)

	}

	err = json.Unmarshal(tags, &note.Tags)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON : %w", err)
	}

	return &note, nil
}
