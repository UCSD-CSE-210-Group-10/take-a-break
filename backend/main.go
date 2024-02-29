// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"fmt"
	"net/http"
	"os"
	"take-a-break/web-service/database"
	"take-a-break/web-service/events"

	"github.com/gin-contrib/cors"

	"take-a-break/web-service/login"

	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // allow request from http://localhost:3000
	router.Use(cors.New(config))

	router.GET("/events", func(c *gin.Context) {
		events.GetEvents(c, conn)
	})
	router.GET("/events/:id", func(c *gin.Context) {
		events.GetEventByID(c, conn)
	})
	router.POST("/events", func(c *gin.Context) {
		events.PostEvent(c, conn)
	})

	router.GET("/GoogleLogin", login.HandleGoogleLogin)
	router.GET("/GoogleCallback", login.HandleGoogleCallback)

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
