package api

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internal/handlers"
	"mobin.dev/internal/service"
)

func RegisterEventsRouter(r *gin.Engine, e service.EventService) {
	h := handlers.NewEventHandler(e)

	eventsGroup := r.Group("/events")

	{
		eventsGroup.GET("/", h.GetEvents)
		eventsGroup.GET("/:id", h.GetEvent)
	}

}
