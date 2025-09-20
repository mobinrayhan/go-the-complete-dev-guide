package repository

import "database/sql"

type NotesRepo struct {
	db *sql.DB
}

func NewNotesRepo(db *sql.DB) *NotesRepo {
	return &NotesRepo{
		db,
	}
}

func (r *NotesRepo) FetchNotes() {

}
