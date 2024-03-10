package user_event

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/user_event"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPostUserEvent(t *testing.T) {
	// Create a temporary database connection for testing
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Create a new Gin router
	router := gin.Default()

	// Define the route
	router.POST("/user_event", func(c *gin.Context) {
		user_event.PostUserEvent(c, conn)
	})

	// Create a sample user data for testing
	userEventData := models.UserEvent{
		EmailID: "user1@example.com",
		EventID: "1",
	}

	// Convert user data to JSON
	jsonData, err := json.Marshal(userEventData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the JSON data
	req := httptest.NewRequest("POST", "/user_event", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder
	recorder := httptest.NewRecorder()

	// Handle the request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// Clean up the test data
	_, err = conn.ExecuteQuery("DELETE FROM user_event WHERE email_id = $1", userEventData.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
}

func TestGetUserEvent(t *testing.T) {
	// Create a temporary database connection for testing
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Create a new Gin router
	router := gin.Default()

	// Define the route
	router.GET("/user_event/:email_id/:event_id", func(c *gin.Context) {
		user_event.GetUserEvent(c, conn)
	})

	// Insert a sample user event into the database for testing
	_, err = conn.ExecuteQuery("INSERT INTO user_event (email_id, event_id) VALUES ($1, $2)", "user1@example.com", 1)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/user_event/user1@example.com/1", nil)

	// Create a new HTTP recorder
	recorder := httptest.NewRecorder()

	// Handle the request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	var response models.UserEvent
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "user1@example.com", response.EmailID)
	assert.Equal(t, "1", response.EventID)
}
