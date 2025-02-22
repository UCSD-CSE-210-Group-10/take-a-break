// reference: https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/events"
	"take-a-break/web-service/friend_request"
	"take-a-break/web-service/friends"
	"take-a-break/web-service/login"
	"take-a-break/web-service/user_event"
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
	config.AllowOrigins = []string{os.Getenv("CLIENT_URL")}
	router.Use(cors.New(config))

	router.GET("/events/all/:token", func(c *gin.Context) {
		events.GetEvents(c, conn)
	})
	router.GET("/events/:id", func(c *gin.Context) {
		events.GetEventByID(c, conn)
	})
	router.POST("/events", func(c *gin.Context) {
		events.PostEvent(c, conn)
	})

	router.GET("/events/search", func(c *gin.Context) {
		events.SearchEvents(c, conn)
	})

	// router.POST("/users", func(c *gin.Context) {
	// 	users.PostUser(c, conn)
	// })

	router.GET("/friends/:token", func(c *gin.Context) {
		friends.GetFriendsByEmailID(c, conn)
	})

	router.GET("/users/:token", func(c *gin.Context) {
		users.GetUserByEmailID(c, conn)
	})

	router.POST("/friends/request/send/:token", func(c *gin.Context) {
		friend_request.PostFriendRequest(c, conn)
	})

	router.POST("/friends/request/accept/:token", func(c *gin.Context) {
		friend_request.PostAcceptFriendRequest(c, conn)
	})

	router.POST("/friends/request/ignore/:token", func(c *gin.Context) {
		friend_request.PostIgnoreFriendRequest(c, conn)
	})
	router.POST("/user_event/:token/:event_id", func(c *gin.Context) {
		user_event.PostUserEvent(c, conn)
	})
	router.GET("/user_event/:token/:event_id", func(c *gin.Context) {
		user_event.GetUserEvent(c, conn)
	})

	router.GET("/friends/attendance/:token/:id", user_event.GetFriendsAttendingEventHandler(conn))

	router.GET("/friends/request/get/:token", func(c *gin.Context) {
		friend_request.GetFriendRequests(c, conn)
	})

	router.GET("/friends/search/:token", friends.SearchFriendsHandler(conn))
	// router.POST("/delete-friend", handle_friend.DeleteFriendHandler(conn))

	router.GET("/auth/token", func(c *gin.Context) {
		login.GetLoginHandler(c, conn)
	})

	router.GET("/auth/verify/:token", auth.GetAuthTokenHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
