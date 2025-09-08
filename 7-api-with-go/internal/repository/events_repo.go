package repository

import (
	"time"

	"mobin.dev/internal/models"
)

type EventsRepo struct{}

func NewEventsRepo() *EventsRepo {
	return &EventsRepo{}
}

func (r *EventsRepo) FetchAllEvents() []models.Event {
	return []models.Event{
		{ID: "1", Name: "Go Meetup", Location: "Dhaka", Date: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Name: "Gin Workshop", Location: "Chittagong", Date: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
}

func (r *EventsRepo) FetchEvent() models.Event {
	return models.Event{
		ID:       "2",
		Name:     "Gin Workshop",
		Location: "Chittagong", Date: time.Now(),
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

}
