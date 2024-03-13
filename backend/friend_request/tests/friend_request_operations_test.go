package friend_request

import (
	"fmt"
	"take-a-break/web-service/database"
	"take-a-break/web-service/friend_request"
	"take-a-break/web-service/friends"
	"take-a-break/web-service/models"
	"take-a-break/web-service/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendFriendRequest(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	user1 := models.User{
		EmailID: "test_user1@example.com",
		Name:    "Test User 1",
		Role:    "user",
		Avatar:  "test-avatar1",
	}

	user2 := models.User{
		EmailID: "test_user2@example.com",
		Name:    "Test User 2",
		Role:    "user",
		Avatar:  "test-avatar2",
	}

	// Insert the users into the database
	_, err = users.InsertUserIntoDatabase(conn, user1)
	assert.NoError(t, err, "Failed to insert the user into the database")
	_, err = users.InsertUserIntoDatabase(conn, user2)
	assert.NoError(t, err, "Failed to insert the user into the database")

	// Send friend request
	err = friend_request.SendFriendRequest(conn, user1.EmailID, user2.EmailID)
	assert.NoError(t, err, "Failed to send friend request")

	// Fetch friend requests
	friendRequests, err := friend_request.FetchFriendRequest(conn, user2.EmailID)
	assert.NoError(t, err, "Failed to fetch friend requests")

	// Assert the number of friend requests
	assert.Equal(t, 1, len(friendRequests), "Incorrect number of friend requests")
	assert.Equal(t, user1.Name, friendRequests[0].Name, "Friend request sender name does not match")

	// Clean up
	rows, err := conn.ExecuteQuery("DELETE FROM friend_requests WHERE sender = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()
}

func TestAcceptFriendRequest(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	user1 := models.User{
		EmailID: "test_user1@example.com",
		Name:    "Test User 1",
		Role:    "user",
		Avatar:  "test-avatar1",
	}

	user2 := models.User{
		EmailID: "test_user2@example.com",
		Name:    "Test User 2",
		Role:    "user",
		Avatar:  "test-avatar2",
	}

	// Insert the users into the database
	_, err = users.InsertUserIntoDatabase(conn, user1)
	assert.NoError(t, err, "Failed to insert the user into the database")
	_, err = users.InsertUserIntoDatabase(conn, user2)
	assert.NoError(t, err, "Failed to insert the user into the database")

	// Send friend request
	err = friend_request.SendFriendRequest(conn, user1.EmailID, user2.EmailID)
	assert.NoError(t, err, "Failed to send friend request")

	// Accept friend request
	err = friend_request.AcceptFriendRequest(conn, user1.EmailID, user2.EmailID)
	assert.NoError(t, err, "Failed to accept friend request")

	// Fetch friends for user 1
	curFriends, err := friends.FetchFriends(conn, user1.EmailID)
	assert.NoError(t, err, "Failed to fetch friends")

	// Assert the number of friends
	assert.Equal(t, 1, len(curFriends), "Incorrect number of friends")
	assert.Equal(t, user2.Name, curFriends[0].Name, "Friend name does not match")

	// Clean up
	rows, err := conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1 or email_id2 = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()
}

func TestIgnoreFriendRequest(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	user1 := models.User{
		EmailID: "test_user1@example.com",
		Name:    "Test User 1",
		Role:    "user",
		Avatar:  "test-avatar1",
	}

	user2 := models.User{
		EmailID: "test_user2@example.com",
		Name:    "Test User 2",
		Role:    "user",
		Avatar:  "test-avatar2",
	}

	// Insert the users into the database
	_, err = users.InsertUserIntoDatabase(conn, user1)
	assert.NoError(t, err, "Failed to insert the user into the database")
	_, err = users.InsertUserIntoDatabase(conn, user2)
	assert.NoError(t, err, "Failed to insert the user into the database")

	// Send friend request
	err = friend_request.SendFriendRequest(conn, user1.EmailID, user2.EmailID)
	assert.NoError(t, err, "Failed to send friend request")

	// Ignore friend request
	err = friend_request.IgnoreFriendRequest(conn, user2.EmailID, user1.EmailID)
	assert.NoError(t, err, "Failed to ignore friend request")

	// Fetch friend requests for user 1
	friendRequests, err := friend_request.FetchFriendRequest(conn, user1.EmailID)
	assert.NoError(t, err, "Failed to fetch friend requests")

	// Assert the number of friend requests
	assert.Equal(t, 0, len(friendRequests), "Incorrect number of friend requests")

	// Clean up
	rows, err := conn.ExecuteQuery("DELETE FROM friend_requests WHERE sender = $1 OR reciever = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()
}

func TestFetchFriendRequest(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	user1 := models.User{
		EmailID: "test_user1@example.com",
		Name:    "Test User 1",
		Role:    "user",
		Avatar:  "test-avatar1",
	}

	user2 := models.User{
		EmailID: "test_user2@example.com",
		Name:    "Test User 2",
		Role:    "user",
		Avatar:  "test-avatar2",
	}

	// Insert the users into the database
	_, err = users.InsertUserIntoDatabase(conn, user1)
	assert.NoError(t, err, "Failed to insert the user into the database")
	_, err = users.InsertUserIntoDatabase(conn, user2)
	assert.NoError(t, err, "Failed to insert the user into the database")

	// Send friend request
	err = friend_request.SendFriendRequest(conn, user1.EmailID, user2.EmailID)
	assert.NoError(t, err, "Failed to send friend request")

	// Fetch friend requests for user 2
	friendRequests, err := friend_request.FetchFriendRequest(conn, user2.EmailID)
	assert.NoError(t, err, "Failed to fetch friend requests")

	// Assert the number of friend requests
	assert.Equal(t, 1, len(friendRequests), "Incorrect number of friend requests")
	assert.Equal(t, user1.Name, friendRequests[0].Name, "Friend request sender name does not match")

	// Clean up
	rows, err := conn.ExecuteQuery("DELETE FROM friend_requests WHERE sender = $1 OR reciever = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()
}
