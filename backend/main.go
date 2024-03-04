// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"take-a-break/web-service/database"
	"take-a-break/web-service/events"
	"take-a-break/web-service/users"

	"take-a-break/web-service/handle_friend"
	"take-a-break/web-service/login"

	"github.com/gin-contrib/cors"

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
	router.POST("/users", func(c *gin.Context) {
		users.PostUser(c, conn)
	})
	router.GET("/users/:email_id", func(c *gin.Context) {
		users.GetUserByEmailID(c, conn)
	})
	router.POST("/makefriends", func(c *gin.Context) {
		users.PostFriends(c, conn)
	})

	router.GET("/GoogleLogin", login.HandleGoogleLogin)
	router.GET("/GoogleCallback", login.HandleGoogleCallback)

	router.GET("/login", func(c *gin.Context) {
		// URL for return to login page
		c.JSON(http.StatusOK, gin.H{
			"url": "http://localhost:3000",
		})
	})

	// Example search term
	searchTerm := "Regular User 1"

	// Search for friends
	foundUsers, err := handle_friend.SearchFriends(conn, searchTerm)
	if err != nil {
		log.Fatal("Error searching for friends:", err)
	}

	// Display results
	fmt.Printf("Found %d users:\n", len(foundUsers))
	for _, user := range foundUsers {
		fmt.Printf("Email_id: %s, Name: %s\n", user.Email_id, user.Name)
	}

	http.HandleFunc("/delete-friend", func(w http.ResponseWriter, r *http.Request) {
		emailID1 := r.FormValue("email_id1")
		emailID2 := r.FormValue("email_id2")

		err := handle_friend.DeleteFriend(conn, emailID1, emailID2)
		if err != nil {
			http.Error(w, "Error deleting friend", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Friendship between '%s' and '%s' deleted successfully", emailID1, emailID2)
	})

	http.HandleFunc("/add-friend", func(w http.ResponseWriter, r *http.Request) {
		emailID1 := r.FormValue("email_id1")
		emailID2 := r.FormValue("email_id2")

		err := handle_friend.AddFriend(conn, emailID1, emailID2)
		if err != nil {
			http.Error(w, "Error adding friend", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Friendship between '%s' and '%s' added successfully", emailID1, emailID2)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)

}
