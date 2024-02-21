// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (

	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"take-a-break/web-service/events"


	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // allow request from http://localhost:3000
	router.Use(cors.New(config))

	// router.GET("/events", getEvents)
	// router.GET("/events/:id", getEventByID)
	// router.POST("/events", postEvent)

	router.GET("/GoogleLogin", handleGoogleLogin)
	router.GET("/GoogleCallback", handleGoogleCallback)

	router.GET("/login", func(c *gin.Context) {
		// URL for return to login page
		c.JSON(http.StatusOK, gin.H{
			"url": "http://localhost:3000",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)

}
