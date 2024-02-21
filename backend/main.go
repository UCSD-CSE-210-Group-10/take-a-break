// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"take-a-break/web-service/events"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/events", events.GetEvents)
	router.GET("/events/:id", events.GetEventByID)
	router.POST("/events", events.PostEvent)

	router.Run("localhost:8080")
}
