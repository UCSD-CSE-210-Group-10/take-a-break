package user_event

import (
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/users"

	"github.com/gin-gonic/gin"
)

type User = models.User

// GetFriendsAttendingEventByID retrieves a list of friends attending a specific event
func GetFriendsAttendingEventByID(c *gin.Context, conn *database.DBConnection) ([]User, error) {
	// Get the event ID from the request parameters
	// emailID := c.Param("email_id")
	eventID := c.Param("event_id")

	// Fetch the list of friends for the current user
	// currentUserEmail := emailID // You need to implement this function to get the current user's email
	friends, err := users.GetFriendsByEmailID(c, conn)
	if err != nil {
		return nil, err
	}

	// Initialize a slice to store friends attending the event
	var friendsAttendingEvent []User

	// Check if each friend is attending the event
	for _, friend := range friends {
		isAttending, err := isUserAttendingEvent(conn, friend.EmailID, eventID)
		if err != nil {
			return nil, err
		}
		if isAttending {
			friendsAttendingEvent = append(friendsAttendingEvent, friend)
		}
	}

	// Return the list of friends attending the event
	return friendsAttendingEvent, nil
}

// isUserAttendingEvent checks if the user is attending the specified event
func isUserAttendingEvent(conn *database.DBConnection, userEmail string, eventID string) (bool, error) {
	// Query to check if the user is attending the event
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM user_event
            WHERE email_id = $1 AND event_id = $2
        )
    `

	// Execute the query
	rows, _ := conn.ExecuteQuery(query, userEmail, eventID)

	// Scan the result
	var isAttending bool
	err := rows.Scan(&isAttending)
	if err != nil {
		return false, err
	}

	return isAttending, nil
}
