package dto

import (
	"time"
)

type EventResponse struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

type CreateEventRequest struct {
	Name     string    `json:"name" binding:"required"`
	Location string    `json:"location" binding:"required"`
	Date     time.Time `json:"date" binding:"required"`
}
