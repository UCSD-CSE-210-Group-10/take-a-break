package user_event

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/user_event"
	"take-a-break/web-service/users"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFetchFriendsAttendingEvent(t *testing.T) {
	// Create a new database connection
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Create a new Gin router
	router := gin.Default()

	// Define the route
	router.GET("/friend_attendance/:email_id/:event_id", func(c *gin.Context) {
		user_event.GetFriendsAttendingEventByID(c, conn)
	})

	// Insert some test users into the database
	testUsers := []users.User{
		{EmailID: "friend1@example.com", Name: "Friend 1", Role: "user"},
		{EmailID: "friend2@example.com", Name: "Friend 2", Role: "user"},
		{EmailID: "friend3@example.com", Name: "Friend 3", Role: "user"},
	}

	// Clean up the test data from the database
	for _, user := range testUsers {
		_, err := conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user.EmailID)
		if err != nil {
			t.Fatal(err)
		}
	}
	_, err = conn.ExecuteQuery("DELETE FROM user_event WHERE event_id IN ($1, $2)", "1", "2")
	if err != nil {
		t.Fatal(err)
	}

	for _, user := range testUsers {
		_, err := users.InsertUserIntoDatabase(conn, user)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Insert some test user events into the database
	testUserEvents := []models.UserEvent{
		{EmailID: "friend1@example.com", EventID: "1"},
		{EmailID: "friend2@example.com", EventID: "1"},
		{EmailID: "friend3@example.com", EventID: "2"},
	}
	for _, userEvent := range testUserEvents {
		_, err := user_event.InsertUserEventIntoDatabase(conn, userEvent)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/friend_attendance/user1@example.com/1", nil)

	// Create a new HTTP recorder
	recorder := httptest.NewRecorder()

	// Handle the request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []models.User
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the number of friends attending the event
	assert.Equal(t, 2, len(response), "Incorrect number of friends attending the event")

	// Assert the details of the first friend
	assert.Equal(t, "friend1@example.com", response[0].EmailID, "Incorrect email ID for first friend")
	assert.Equal(t, "Friend 1", response[0].Name, "Incorrect name for first friend")

	// Clean up the test data from the database
	for _, user := range testUsers {
		_, err := conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user.EmailID)
		if err != nil {
			t.Fatal(err)
		}
	}
	_, err = conn.ExecuteQuery("DELETE FROM user_event WHERE event_id IN ($1, $2)", "1", "2")
	if err != nil {
		t.Fatal(err)
	}
}
