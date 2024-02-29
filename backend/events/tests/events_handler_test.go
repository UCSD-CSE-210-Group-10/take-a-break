package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"take-a-break/web-service/database"
	"take-a-break/web-service/events"
	"take-a-break/web-service/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetEvents(t *testing.T) {
	r := SetUpRouter()

	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	r.GET("/events", func(c *gin.Context) {
		events.GetEvents(c, conn)
	})

	req, _ := http.NewRequest("GET", "/events", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var events []models.Event
	json.Unmarshal(w.Body.Bytes(), &events)
	var expectedEvents []models.Event = []models.Event{
		{
			ID:          "1",
			Title:       "Event 1",
			Venue:       "Venue 1",
			Date:        "2024-02-17T00:00:00Z",
			Time:        "0000-01-01T18:00:00Z",
			Description: "Description for Event 1",
			Tags:        "Tag1, Tag2",
			ImagePath:   "./images/event1.jpg",
			Host:        "Host 1",
			Contact:     "Contact 1",
		},
		{
			ID:          "2",
			Title:       "Event 2",
			Venue:       "Venue 2",
			Date:        "2024-02-18T00:00:00Z",
			Time:        "0000-01-01T19:30:00Z",
			Description: "Description for Event 2",
			Tags:        "Tag2, Tag3",
			ImagePath:   "./images/event2.jpg",
			Host:        "Host 2",
			Contact:     "Contact 2",
		},
		{
			ID:          "3",
			Title:       "Event 3",
			Venue:       "Venue 3",
			Date:        "2024-02-19T00:00:00Z",
			Time:        "0000-01-01T20:15:00Z",
			Description: "Description for Event 3",
			Tags:        "Tag3, Tag4",
			ImagePath:   "./images/event3.jpg",
			Host:        "Host 3",
			Contact:     "Contact 3",
		},
	}

	assert.Equal(t, events, expectedEvents)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, events)
}

func TestGetEventById(t *testing.T) {
	r := SetUpRouter()

	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	r.GET("/events/:id", func(c *gin.Context) {
		events.GetEventByID(c, conn)
	})

	req, _ := http.NewRequest("GET", "/events/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var event models.Event
	json.Unmarshal(w.Body.Bytes(), &event)
	var expectedEvent models.Event = models.Event{
		ID:          "1",
		Title:       "Event 1",
		Venue:       "Venue 1",
		Date:        "2024-02-17T00:00:00Z",
		Time:        "0000-01-01T18:00:00Z",
		Description: "Description for Event 1",
		Tags:        "Tag1, Tag2",
		ImagePath:   "./images/event1.jpg",
		Host:        "Host 1",
		Contact:     "Contact 1",
	}

	assert.Equal(t, event, expectedEvent)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, event)
}

func TestPostEvent(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_image.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write some content to the temporary file
	content := []byte("test image content")
	_, err = tempFile.Write(content)
	if err != nil {
		t.Fatal(err)
	}
	tempFile.Close()

	// Create a new multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add form fields
	_ = writer.WriteField("title", "Test Event")
	_ = writer.WriteField("venue", "Test Venue")
	_ = writer.WriteField("date", "2022-12-31")
	_ = writer.WriteField("time", "18:00:00")
	_ = writer.WriteField("description", "Test Description")
	_ = writer.WriteField("tags", "tag1,tag2")
	_ = writer.WriteField("host", "Test Host")
	_ = writer.WriteField("contact", "Test Contact")

	// Create a new file part and add it to the form
	fileWriter, _ := writer.CreateFormFile("image", filepath.Base(tempFile.Name()))
	_, _ = fileWriter.Write(content)

	// Close the multipart writer
	writer.Close()

	// Create a new HTTP request with the form data
	req := httptest.NewRequest("POST", "/events", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a new HTTP recorder
	recorder := httptest.NewRecorder()

	// Create a new Gin router and handle the request
	router := gin.Default()
	router.POST("/events", func(c *gin.Context) {
		events.PostEvent(c, conn) // Replace nil with your database connection
	})
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the expected values
	assert.Equal(t, "Test Event", response["title"])
	assert.Equal(t, "Test Venue", response["venue"])
	// Add assertions for other fields

	// // Assert that the image file was saved
	// // Replace "expectedImagePath" with the expected path of the saved image
	// fmt.Println(response["filename"])
	// expectedImagePath := "./images/test_image.jpg"
	// assert.Equal(t, expectedImagePath, response["filename"])

	// clean up
	_, err = conn.ExecuteQuery("DELETE FROM events WHERE title = $1", "Test Event")
	assert.NoError(t, err, "Failed to clean up the test data")
}
