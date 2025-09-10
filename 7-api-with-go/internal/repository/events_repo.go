package repository

import (
	"fmt"
	"time"

	"mobin.dev/internal/models"
)

type EventsRepo struct {
	Value string
}

func NewEventsRepo(value string) *EventsRepo {
	return &EventsRepo{
		Value: value,
	}
}

func (r *EventsRepo) FetchAllEvents() []models.Event {
	fmt.Println(r.Value)
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
