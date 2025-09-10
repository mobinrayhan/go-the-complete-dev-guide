package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internal/dto"
	"mobin.dev/internal/service"
)

type EventsHandlers struct {
	Service service.EventService
}

func NewEventHandler(s service.EventService) EventsHandlers {
	return EventsHandlers{Service: s}
}

func (h EventsHandlers) GetEvents(context *gin.Context) {
	events := h.Service.GetAllEvents()

	var resp []dto.EventResponse

	for _, e := range events {
		resp = append(resp, dto.EventResponse{
			ID:        e.ID,
			Name:      e.Name,
			Location:  e.Location,
			Date:      e.Date,
			CreatedAt: e.CreatedAt,
		})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Fetched All Events Successfully", "events": resp})
}

func (h EventsHandlers) GetEvent(context *gin.Context) {
	event := h.Service.GetEvent()

	eventResp := dto.EventResponse{
		ID:       event.ID,
		Name:     event.Name,
		Location: event.Location,
		Date:     event.Date,
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Fetched Single User Successfully!", "event": eventResp})
}
