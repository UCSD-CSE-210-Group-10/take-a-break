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

	testUser := models.User{
		EmailID: "testuser@example.com",
		Name:    "Test User",
		Role:    "user",
		Avatar:  "test-avatar",
	}

	userEvent := models.UserEvent{
		EmailID: "testuser@example.com",
		EventID: "1",
	}

	// Insert a sample user into the database for testing
	rows, err := conn.ExecuteQuery("INSERT INTO users (email_id, name, role, avatar) VALUES ($1, $2, $3, $4)", testUser.EmailID, testUser.Name, testUser.Role, testUser.Avatar)
	assert.NoError(t, err, "Failed to insert the test user into users table")
	defer rows.Close()

	insertedUserEvent, err := user_event.InsertUserEventIntoDatabase(conn, userEvent.EmailID, userEvent.EventID)
	assert.NoError(t, err, "Failed to insert user event into database")

	assert.Equal(t, userEvent.EmailID, insertedUserEvent.EmailID, "Email ID does not match")
	assert.Equal(t, userEvent.EventID, insertedUserEvent.EventID, "Event ID does not match")

	// Clean up
	rows, err = conn.ExecuteQuery("DELETE FROM user_event WHERE email_id = $1", userEvent.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()
	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", testUser.EmailID)
	assert.NoError(t, err, "Failed to clean up test user data")
	defer rows.Close()
}

func TestGetUserEventFromDatabase(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	testUser := models.User{
		EmailID: "testuser@example.com",
		Name:    "Test User",
		Role:    "user",
		Avatar:  "test-avatar",
	}

	userEvent := models.UserEvent{
		EmailID: "testuser@example.com",
		EventID: "1",
	}

	// Insert a sample user into the database for testing
	rows, err := conn.ExecuteQuery("INSERT INTO users (email_id, name, role, avatar) VALUES ($1, $2, $3, $4)", testUser.EmailID, testUser.Name, testUser.Role, testUser.Avatar)
	assert.NoError(t, err, "Failed to insert the test user into users table")
	defer rows.Close()
	// Insert the sample user event into the database
	rows, err = conn.ExecuteQuery("INSERT INTO user_event (email_id, event_id) VALUES ($1, $2)", userEvent.EmailID, userEvent.EventID)
	assert.NoError(t, err, "Failed to insert the test data into user_event table")
	defer rows.Close()

	// Retrieve the user event from the database
	retrievedUserEvent, err := user_event.GetUserEventFromDatabase(conn, userEvent.EmailID, userEvent.EventID)

	// Check if there were any errors during retrieval
	assert.NoError(t, err, "Failed to retrieve user event from database")

	// Check if the retrieved user event matches the expected user event
	assert.Equal(t, userEvent.EmailID, retrievedUserEvent.EmailID, "Email ID does not match")
	assert.Equal(t, userEvent.EventID, retrievedUserEvent.EventID, "Event ID does not match")

	// Clean up
	rows, err = conn.ExecuteQuery("DELETE FROM user_event WHERE email_id = $1", userEvent.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", testUser.EmailID)
	assert.NoError(t, err, "Failed to clean up test user data")
	defer rows.Close()
}
