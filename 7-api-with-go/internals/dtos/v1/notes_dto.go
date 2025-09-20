package dtosV1

import "time"

type NoteResponse struct {
	Id        int         `json:"id"`
	UserId    int         `json:"userId"`
	Title     string      `json:"title"`
	Body      string      `json:"body"`
	Tags      interface{} `json:"tags"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}
