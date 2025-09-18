package model

import "time"

type Note struct {
	Id        int
	UserId    int
	Title     string
	Body      string
	Tags      interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
}
