package user_event

import (
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/user_event"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertUserEventIntoDatabase(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	userEvent := models.UserEvent{
		EmailID: "user1@example.com",
		EventID: "1",
	}

	insertedUserEvent, err := user_event.InsertUserEventIntoDatabase(conn, userEvent)

	// Check if there were any errors during insertion
	assert.NoError(t, err, "Failed to insert user event into database")

	assert.Equal(t, userEvent.EmailID, insertedUserEvent.EmailID, "Email ID does not match")
	assert.Equal(t, userEvent.EventID, insertedUserEvent.EventID, "Event ID does not match")

	// Clean up
	_, err = conn.ExecuteQuery("DELETE FROM user_event WHERE email_id = $1", userEvent.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
}

func TestGetUserEventFromDatabase(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Define a sample user event
	emailID := "user1@example.com"
	eventID := "1"

	// Insert the sample user event into the database
	_, err = conn.ExecuteQuery("INSERT INTO user_event (email_id, event_id) VALUES ($1, $2)", emailID, eventID)
	if err != nil {
		t.Fatal(err)
	}

	// Retrieve the user event from the database
	retrievedUserEvent, err := user_event.GetUserEventFromDatabase(conn, emailID, eventID)

	// Check if there were any errors during retrieval
	assert.NoError(t, err, "Failed to retrieve user event from database")

	// Check if the retrieved user event matches the expected user event
	assert.Equal(t, emailID, retrievedUserEvent.EmailID, "Email ID does not match")
	assert.Equal(t, eventID, retrievedUserEvent.EventID, "Event ID does not match")

	// Clean up
	_, err = conn.ExecuteQuery("DELETE FROM user_event WHERE email_id = $1", emailID)
	assert.NoError(t, err, "Failed to clean up the test data")
}
