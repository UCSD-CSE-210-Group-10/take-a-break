package events

import (
	"fmt"
	"take-a-break/web-service/database"
	"take-a-break/web-service/events"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertEventIntoDatabase(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Create a sample event data
	formData := map[string]string{
		"title":       "Sample Event",
		"description": "This is a sample event",
		"date":        "2021-12-31",
		"time":        "12:00",
	}

	event, err := events.InsertEventIntoDatabase(conn, formData)

	assert.NoError(t, err, "Failed to insert the event into the database")
	assert.Equal(t, event.Title, formData["title"], "The inserted event title does not match the provided title")
	assert.Equal(t, event.Description, formData["description"], "The inserted event description does not match the provided description")

	// Clean up
	_, err = conn.ExecuteQuery("DELETE FROM events WHERE title = $1", formData["title"])
	assert.NoError(t, err, "Failed to clean up the test data")
}

func TestFetchEventByID(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ID := "1"
	fetchedEvent, err := events.FetchEventByID(conn, ID)

	assert.NoError(t, err, "Failed to fetch the event from the database")
	assert.Equal(t, fetchedEvent.ID, ID, "The fetched event ID does not match the inserted event ID")
}

func TestFetchAllEvents(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	LEN_EVENTS := 3

	fetchedEvents, err := events.FetchAllEvents(conn)

	assert.NoError(t, err, "Failed to fetch the events from the database")
	assert.Equal(t, len(fetchedEvents), LEN_EVENTS, "FetchAllEvents did not fetch all the events")
}
