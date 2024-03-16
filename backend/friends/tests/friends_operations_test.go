package friends

import (
	"fmt"
	"take-a-break/web-service/database"
	"take-a-break/web-service/friends"
	"take-a-break/web-service/models"
	"take-a-break/web-service/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeFriends(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	user1 := models.User{
		EmailID: "test_user9@example.com",
		Name:    "Test User 9",
		Role:    "user",
		Avatar:  "test-avatar9",
	}

	user2 := models.User{
		EmailID: "test_user10@example.com",
		Name:    "Test User 10",
		Role:    "user",
		Avatar:  "test-avatar10",
	}

	// Insert the users into the database
	_, err = users.InsertUserIntoDatabase(conn, user1)
	assert.NoError(t, err, "Failed to insert the user into the database")
	_, err = users.InsertUserIntoDatabase(conn, user2)
	assert.NoError(t, err, "Failed to insert the user into the database")

	// Make friends
	_, err = friends.MakeFriends(conn, user1.EmailID, user2.EmailID)
	assert.NoError(t, err, "Failed to make friends")

	// Fetch friends
	cur_friends, err := friends.FetchFriends(conn, user1.EmailID)
	assert.NoError(t, err, "Failed to fetch friends")

	// Assert the number of friends
	assert.Equal(t, 1, len(cur_friends), "Incorrect number of friends")
	assert.Equal(t, user2.Name, cur_friends[0].Name, "Friend name does not match")

	// Fetch friends for user 2
	cur_friends, err = friends.FetchFriends(conn, user2.EmailID)
	assert.NoError(t, err, "Failed to fetch friends")

	// Assert the number of friends
	assert.Equal(t, 1, len(cur_friends), "Incorrect number of friends")
	assert.Equal(t, user1.Name, cur_friends[0].Name, "Friend name does not match")

	// Clean up
	rows, err := conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1 or email_id2 = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1 or email_id2 = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()

	rows, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()
}

func TestFetchFriends(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Fetch friends
	EMAIL_ID := "user1@example.com"
	cur_friends, err := friends.FetchFriends(conn, EMAIL_ID)
	assert.NoError(t, err, "Failed to fetch friends")

	// Assert the number of friends
	assert.Equal(t, 2, len(cur_friends), "Incorrect number of friends")
	assert.Equal(t, "Admin User", cur_friends[0].Name, "Friend name does not match")
	assert.Equal(t, "Regular User 2", cur_friends[1].Name, "Friend name does not match")
}

func TestSearchFriends(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	search_res, err := friends.SearchFriends(conn, "User", "user1@example.com")
	assert.NoError(t, err, "Failed to fetch friends")

	assert.Equal(t, 3, len(search_res), "Incorrect number of Search Results")
}
