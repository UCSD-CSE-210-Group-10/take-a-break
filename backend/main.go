// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	//"take-a-break/web-service/events"
	"take-a-break/web-service/login"
	"take-a-break/web-service/search_friend"
	"take-a-break/web-service/database"

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

	router.GET("/GoogleLogin", login.HandleGoogleLogin)
	router.GET("/GoogleCallback", login.HandleGoogleCallback)

	router.GET("/login", func(c *gin.Context) {
		// URL for return to login page
		c.JSON(http.StatusOK, gin.H{
			"url": "http://localhost:3000",
		})
	})

	// Create a new database connection
    conn, err := database.NewDBConnection()
    if err != nil {
        log.Fatal("Error establishing database connection:", err)
    }
    defer conn.Close()

    // Example search term
    searchTerm := "john"

    // Search for friends
    foundUsers, err := conn.SearchFriends(searchTerm)
    if err != nil {
        log.Fatal("Error searching for friends:", err)
    }

    // Display results
    fmt.Printf("Found %d users:\n", len(foundUsers))
    for _, user := range foundUsers {
        fmt.Printf("Username: %s, Name: %s\n", user.Username, user.Name)
    }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)

}
