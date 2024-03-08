// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"take-a-break/web-service/database"
	"take-a-break/web-service/events"
	"take-a-break/web-service/handle_friend"
	"take-a-break/web-service/login"
	"take-a-break/web-service/users"

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

	router.GET("/search-friends", handle_friend.SearchFriendsHandler(conn))
	router.POST("/delete-friend", handle_friend.DeleteFriendHandler(conn))

	router.GET("/auth/token", login.GetAuthTokenHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
