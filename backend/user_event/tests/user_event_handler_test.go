package user_event

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"take-a-break/web-service/database"
// 	"take-a-break/web-service/models"
// 	"take-a-break/web-service/user_event"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	return router
// }

// func TestPostUserEvent(t *testing.T) {
// 	// Create a temporary database connection for testing
// 	conn, err := database.NewDBConnection()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Define the route
// 	router.POST("/user_event/:token/:event_id", func(c *gin.Context) {
// 		user_event.PostUserEvent(c, conn)
// 	})

// 	test_token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjA4YmY1YzM3NzJkZDRlN2E3MjdhMTAxYmY1MjBmNjU3NWNhYzMyNmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTgwNTU4MDg2MDY2NjUwNzI5MzAiLCJoZCI6InVjc2QuZWR1IiwiZW1haWwiOiJhYnVkaGlyYWphQHVjc2QuZWR1IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJ6Q1BWU2RPYnI4NVczT0dRYjlNdDBRIiwibmFtZSI6IkFubW9sIEJ1ZGhpcmFqYSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NKZFVsVkYwMmZoOTBOby1CR3JydVJMOS1rRDFPejNCLTFtM3l0Q19vY1g9czk2LWMiLCJnaXZlbl9uYW1lIjoiQW5tb2wiLCJmYW1pbHlfbmFtZSI6IkJ1ZGhpcmFqYSIsImxvY2FsZSI6ImVuIiwiaWF0IjoxNzEwMDQ3OTE3LCJleHAiOjE3MTAwNTE1MTd9.eBiNMvJNXwhpkf0qA1o6EbVxGX7rfu9AqhHTUFKYaDuU9mfmaMGOVUBFTZXIK6XXsXqV5GInViaF1VsWzmbVePTTwzJD0-u9rzT5XLdhChzLcTEH2JvzDwH1QA63S_fcnkZj3d6ewN7OQ2zNeKCAyFbOiaQAnTARp54R5Qp9x4v9ns04qU76a1WNa_a-NkROpRtLyVIZMLJbs3L0yEsj-Lct3zV2F5YWXqHeU-u3i5Iivnj4nAO3TDzXC_hrQXt8yijFIUZowuXwt1A-VKxsOkP4AvnpGyB5OopEByLFA1TuwxScJrD56FUZTNclSIhUWAA3hsNiajkZ8R9ypLe54Q"

// 	// Insert a sample user into the database for testing
// 	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, 'user')", "abudhiraja@ucsd.edu", "Anmol Budhiraja")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a sample user data for testing
// 	userEventData := models.UserEvent{
// 		EmailID: "abudhiraja@ucsd.edu",
// 		EventID: "1",
// 	}

// 	requestRoute := "/user_event/" + test_token + "/" + userEventData.EventID

// 	// Create a new HTTP request with the JSON data
// 	req := httptest.NewRequest("POST", requestRoute, nil)
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a new HTTP recorder
// 	recorder := httptest.NewRecorder()

// 	// Handle the request
// 	router.ServeHTTP(recorder, req)
// 	fmt.Println("Response Body:", recorder.Body.String())

// 	// Check the response status code
// 	assert.Equal(t, http.StatusCreated, recorder.Code)

// 	// Clean up the test data
// 	_, err = conn.ExecuteQuery("DELETE FROM user_event WHERE email_id = $1", userEventData.EmailID)
// 	assert.NoError(t, err, "Failed to clean up the test data")

// 	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", userEventData.EmailID)
// 	assert.NoError(t, err, "Failed to clean up the test data")

// }

// func TestGetUserEvent(t *testing.T) {
// 	// Create a temporary database connection for testing
// 	conn, err := database.NewDBConnection()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Define the route
// 	router.GET("/user_event/:token/:event_id", func(c *gin.Context) {
// 		user_event.GetUserEvent(c, conn)
// 	})

// 	// Insert a sample user event into the database for testing
// 	_, err = conn.ExecuteQuery("INSERT INTO user_event (email_id, event_id) VALUES ($1, $2)", "user1@example.com", 1)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a new HTTP request
// 	req := httptest.NewRequest("GET", "/user_event/user1@example.com/1", nil)

// 	// Create a new HTTP recorder
// 	recorder := httptest.NewRecorder()

// 	// Handle the request
// 	router.ServeHTTP(recorder, req)

// 	// Check the response status code
// 	assert.Equal(t, http.StatusOK, recorder.Code)

// 	var response models.UserEvent
// 	err = json.Unmarshal(recorder.Body.Bytes(), &response)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, "user1@example.com", response.EmailID)
// 	assert.Equal(t, "1", response.EventID)
// }
