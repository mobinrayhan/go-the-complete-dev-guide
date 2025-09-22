package dtosV1

import "time"

type NoteResponse struct {
	Id        int         `json:"id,omitempty"`
	UserId    int         `json:"userId,omitempty"`
	Title     string      `json:"title,omitempty"`
	Body      string      `json:"body,omitempty"`
	Tags      interface{} `json:"tags,omitempty"`
	CreatedAt time.Time   `json:"createdAt,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
}
