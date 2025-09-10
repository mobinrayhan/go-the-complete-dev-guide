package service

import (
	"mobin.dev/internal/models"
	"mobin.dev/internal/repository"
)

type EventService struct {
	repo repository.EventsRepo
}

func NewEventsService(repo repository.EventsRepo) EventService {
	return EventService{repo: repo}
}

func (e EventService) GetAllEvents() []models.Event {
	allEvents := e.repo.FetchAllEvents()
	return allEvents
}

func (e EventService) GetEvent() models.Event {
	event := e.repo.FetchEvent()
	return event
}
