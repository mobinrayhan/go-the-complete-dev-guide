package v1

import "time"

type CreateNoteRequest struct {
	UserId    int         `json:"userId" binding:"required"`
	Title     string      `json:"title" binding:"required"`
	Body      string      `json:"body" binding:"required"`
	Tags      interface{} `json:"tags"`
	CreatedAt time.Time   `json:"createdAt" binding:"required"`
	UpdatedAt time.Time   `json:"updatedAt" binding:"required"`
}

type NoteResponse struct {
	Id        int         `json:"id"`
	UserId    int         `json:"userId"`
	Title     string      `json:"title"`
	Body      string      `json:"body"`
	Tags      interface{} `json:"tags"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}
