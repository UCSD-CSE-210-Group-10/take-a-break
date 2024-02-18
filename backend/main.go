// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/events", getEvents)
    router.GET("/events/:id", getEventByID)
    router.POST("/events", postEvent)

    router.Run("localhost:8080")
}


