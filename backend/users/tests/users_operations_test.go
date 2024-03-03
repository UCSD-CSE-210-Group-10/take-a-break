package users

import (
	"fmt"
	"take-a-break/web-service/database"
	"take-a-break/web-service/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertUserIntoDatabase(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	user := users.User{
		EmailID: "test_user@example.com",
		Name:    "Test User",
		Role:    "user",
	}

	insertedUser, err := users.InsertUserIntoDatabase(conn, user)

	assert.NoError(t, err, "Failed to insert the user into the database")
	assert.Equal(t, user.EmailID, insertedUser.EmailID, "Email ID does not match")
	assert.Equal(t, user.Name, insertedUser.Name, "Name does not match")

	// Clean up
	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
}

func TestFetchUserByID(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Fetch the user by ID
	USER_EMAIL_ID := "user1@example.com"
	USER_NAME := "Regular User 1"
	USER_ROLE := "user"

	fetchedUser, err := users.FetchUserByEmailID(conn, USER_EMAIL_ID)
	assert.NoError(t, err, "Failed to fetch the user by Email ID")
	assert.Equal(t, USER_EMAIL_ID, fetchedUser.EmailID, "Email ID does not match")
	assert.Equal(t, USER_NAME, fetchedUser.Name, "Name does not match")
	assert.Equal(t, USER_ROLE, fetchedUser.Role, "Role does not match")
}

func TestMakeFriends(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	user1 := users.User{
		EmailID: "test_user1@example.com",
		Name:    "Test User 1",
		Role:    "user",
	}

	user2 := users.User{
		EmailID: "test_user2@example.com",
		Name:    "Test User 2",
		Role:    "user",
	}

	// Insert the users into the database
	_, err = users.InsertUserIntoDatabase(conn, user1)
	assert.NoError(t, err, "Failed to insert the user into the database")
	_, err = users.InsertUserIntoDatabase(conn, user2)
	assert.NoError(t, err, "Failed to insert the user into the database")

	// Make friends
	err = users.MakeFriends(conn, user1.EmailID, user2.EmailID)
	assert.NoError(t, err, "Failed to make friends")

	// Fetch friends
	friends, err := users.FetchFriends(conn, user1.EmailID)
	assert.NoError(t, err, "Failed to fetch friends")

	// Assert the number of friends
	assert.Equal(t, 1, len(friends), "Incorrect number of friends")
	assert.Equal(t, user2.Name, friends[0].Name, "Friend name does not match")

	// Fetch friends for user 2
	friends, err = users.FetchFriends(conn, user2.EmailID)
	assert.NoError(t, err, "Failed to fetch friends")

	// Assert the number of friends
	assert.Equal(t, 1, len(friends), "Incorrect number of friends")
	assert.Equal(t, user1.Name, friends[0].Name, "Friend name does not match")

	// Clean up
	_, err = conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1 or email_id2 = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")

	_, err = conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1 or email_id2 = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")

	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user1.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")

	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user2.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
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
	friends, err := users.FetchFriends(conn, EMAIL_ID)
	assert.NoError(t, err, "Failed to fetch friends")

	// Assert the number of friends
	assert.Equal(t, 2, len(friends), "Incorrect number of friends")
	assert.Equal(t, "Admin User", friends[0].Name, "Friend name does not match")
	assert.Equal(t, "Regular User 2", friends[1].Name, "Friend name does not match")
}
