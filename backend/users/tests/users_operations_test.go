package users

import (
	"fmt"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
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

	user := models.User{
		EmailID: "test_user@example.com",
		Name:    "Test User",
		Role:    "user",
		Avatar:  "test-avatar",
	}

	insertedUser, err := users.InsertUserIntoDatabase(conn, user)

	assert.NoError(t, err, "Failed to insert the user into the database")
	assert.Equal(t, user.EmailID, insertedUser.EmailID, "Email ID does not match")
	assert.Equal(t, user.Name, insertedUser.Name, "Name does not match")

	// Clean up
	rows, err := conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", user.EmailID)
	assert.NoError(t, err, "Failed to clean up the test data")
	defer rows.Close()
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
