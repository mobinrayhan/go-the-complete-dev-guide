package service

import (
	"mobin.dev/internal/models"
	"mobin.dev/internal/repository"
)

var events = repository.NewEventsRepo()

func GetAllEvents() []models.Event {
	events := events.FetchAllEvents()
	return events
}

func GetEvent() models.Event {
	event := events.FetchEvent()
	return event
}
