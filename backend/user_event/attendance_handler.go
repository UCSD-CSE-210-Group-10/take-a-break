package user_event

import (
	"log"
	"net/http"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"

	"github.com/gin-gonic/gin"
)

type User = models.User

// // GetFriendsAttendingEventByID retrieves a list of friends attending a specific event
// func GetFriendsAttendingEventByID(c *gin.Context, conn *database.DBConnection) ([]User, error) {
// 	// Get the event ID from the request parameters
// 	// emailID := c.Param("email_id")
// 	eventID := c.Param("event_id")

// 	// Fetch the list of friends for the current user
// 	// currentUserEmail := emailID // You need to implement this function to get the current user's email
// 	friends, err := users.GetFriendsByEmailID(c, conn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Initialize a slice to store friends attending the event
// 	var friendsAttendingEvent []User

// 	// Check if each friend is attending the event
// 	for _, friend := range friends {

// 		isAttending, err := isUserAttendingEvent(conn, friend.EmailID, eventID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fmt.Printf("Friend %s is attending: %t\n", friend.EmailID, isAttending)
// 		if isAttending {
// 			friendsAttendingEvent = append(friendsAttendingEvent, friend)
// 		}
// 	}
// 	//fmt.Printf("the friendsAttendingEvent is %s \n", friendsAttendingEvent)
// 	// Return the list of friends attending the event
// 	return friendsAttendingEvent, nil
// }

// // isUserAttendingEvent checks if the user is attending the specified event
// func isUserAttendingEvent(conn *database.DBConnection, userEmail string, eventID string) (bool, error) {
// 	// Query to check if the user is attending the event
// 	//fmt.Printf("Checking attendance for user %s at event %s\n", userEmail, eventID)

// 	query := `
//         SELECT EXISTS (
//             SELECT 1
//             FROM user_event
//             WHERE email_id = $1 AND event_id = $2
//         )
//     `

// 	// Execute the query
// 	//rows, err := conn.ExecuteQuery(query, userEmail, eventID)
// 	//fmt.Printf("rows: %+v \n", rows)
// 	// // Scan the result
// 	// var isAttending bool
// 	// err := rows.Scan(&isAttending)
// 	// if err != nil {
// 	// 	return false, err
// 	// }
// 	// fmt.Printf("Checking  %s \n", isAttending)

// 	// return isAttending, nil

// 	// Execute the query and get a single row
// 	var isAttending bool
// 	row, err := conn.QueryRow(query, userEmail, eventID)
// 	if err != nil {
// 		return false, err
// 	}

// 	err = row.Scan(&isAttending)
// 	if err != nil {
// 		return false, err
// 	}

// 	//fmt.Printf("isAttending %t \n", isAttending)

//		return isAttending, nil
//	}
func GetFriendsAttendingEvent(conn *database.DBConnection, emailID, eventID string) ([]User, error) {
	// Your SQL query to get friends attending the specified event
	query := `
	SELECT u.email_id, u.name
	FROM users u
	JOIN user_event ue ON u.email_id = ue.email_id
	WHERE ue.event_id = $2 AND ue.email_id != $1;
    `
	rows, err := conn.ExecuteQuery(query, emailID, eventID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendingFriends []User
	for rows.Next() {
		var friend User
		if err := rows.Scan(&friend.EmailID, &friend.Name); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		attendingFriends = append(attendingFriends, friend)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendingFriends, nil
}
func GetFriendsAttendingEventHandler(conn *database.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID := c.Query("emailID") // Assuming you get emailID as a query parameter
		eventID := c.Query("eventID") // Assuming you get eventID as a query parameter

		friends, err := GetFriendsAttendingEvent(conn, emailID, eventID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching attending friends"})
			return
		}

		c.JSON(http.StatusOK, friends)
	}
}
