package api

import "github.com/gin-gonic/gin"

func RegisterEventsRouter(r *gin.Engine) {

	eventsGroup := r.Group("/events")

	{
		eventsGroup.GET("/", GetEvents)
		eventsGroup.GET("/:id", GetEvent)
	}

}
