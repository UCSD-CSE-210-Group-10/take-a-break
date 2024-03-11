package user_event

import (
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/user_event"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFriendsAttendingEvent(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	testUser := models.User{
		EmailID: "abudhiraja@ucsd.edu",
		Name:    "Anmol Budhiraja",
		Role:    "user",
	}
	testFriend1 := models.User{
		EmailID: "testfriend1@example.com",
		Name:    "Test Friend 1",
		Role:    "user",
	}
	testFriend2 := models.User{
		EmailID: "testfriend2@example.com",
		Name:    "Test Friend 2",
		Role:    "user",
	}

	testEventID := "1"

	// Insert sample users into the database for testing
	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, $3)", testUser.EmailID, testUser.Name, testUser.Role)
	assert.NoError(t, err, "Failed to insert the test user into users table")
	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, $3)", testFriend1.EmailID, testFriend1.Name, testFriend1.Role)
	assert.NoError(t, err, "Failed to insert the test friend 1 into users table")
	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, $3)", testFriend2.EmailID, testFriend2.Name, testFriend2.Role)
	assert.NoError(t, err, "Failed to insert the test friend 2 into users table")

	// Insert sample friends into the database for testing
	_, err = conn.ExecuteQuery("INSERT INTO friends (email_id1, email_id2) VALUES ($1, $2)", testUser.EmailID, testFriend1.EmailID)
	assert.NoError(t, err, "Failed to insert user&friend1 into friends table")
	_, err = conn.ExecuteQuery("INSERT INTO friends (email_id1, email_id2) VALUES ($1, $2)", testUser.EmailID, testFriend2.EmailID)
	assert.NoError(t, err, "Failed to insert user&friend2 into friends table")

	// Insert the sample user event into the database
	_, err = conn.ExecuteQuery("INSERT INTO user_event (email_id, event_id) VALUES ($1, $2)", testFriend1.EmailID, testEventID)
	assert.NoError(t, err, "Failed to insert the test data into user_event table")

	// Retrieve the attendingFriends List from the database
	attendingFriends, err := user_event.GetFriendsAttendingEvent(conn, testUser.EmailID, testEventID)

	assert.NoError(t, err, "Failed to retrieve attending friends from database")
	assert.Equal(t, 1, len(attendingFriends), "friendsAttending length doesn't match")
	assert.Equal(t, testFriend1.EmailID, attendingFriends[0].EmailID, "Email ID does not match")

	// Clean up
	_, err = conn.ExecuteQuery("DELETE FROM user_event WHERE email_id = $1", testFriend1.EmailID)
	assert.NoError(t, err, "Failed to clean up test user_event data")
	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", testUser.EmailID)
	assert.NoError(t, err, "Failed to clean up test user data")
	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", testFriend1.EmailID)
	assert.NoError(t, err, "Failed to clean up test friend1 data")
	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", testFriend2.EmailID)
	assert.NoError(t, err, "Failed to clean up test friend2 data")
	_, err = conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1", testUser.EmailID)
	assert.NoError(t, err, "Failed to clean up test friends data")
}
