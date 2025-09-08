package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internal/api/dto"
	"mobin.dev/internal/service"
)

// var eventsService = service.NewEventsService()

func GetEvents(context *gin.Context) {
	events := service.GetAllEvents()

	var resp []dto.EventResponse

	for _, e := range events {
		resp = append(resp, dto.EventResponse{
			ID:       e.ID,
			Name:     e.Name,
			Location: e.Location,
			Date:     e.Date,
		})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Fetched All Events Successfully", "events": resp})
}

func GetEvent(context *gin.Context) {
	event := service.GetEvent()

	eventResp := dto.EventResponse{
		ID:       event.ID,
		Name:     event.Name,
		Location: event.Location,
		Date:     event.Date,
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Fetched Single User Successfully!", "event": eventResp})
}
