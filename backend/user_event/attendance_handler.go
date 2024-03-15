package user_event

import (
	"log"
	"net/http"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"

	"github.com/gin-gonic/gin"
)

type User = models.User

func GetFriendsAttendingEvent(conn *database.DBConnection, emailID, eventID string) ([]User, error) {
	// Your SQL query to get friends attending the specified event
	query := `
	SELECT u.email_id, u.name, u.avatar
	FROM users u
	JOIN user_event ue ON u.email_id = ue.email_id
	INNER JOIN friends f ON f.email_id1 = $1 AND f.email_id2 = u.email_id
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
		if err := rows.Scan(&friend.EmailID, &friend.Name, &friend.Avatar); err != nil {
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

		token := c.Param("token")

		if !auth.VerifyJWTToken(token) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
			return
		}

		claims := auth.ReturnJWTToken(token)

		emailID := claims["email"].(string)

		eventID := c.Param("id")

		friends, err := GetFriendsAttendingEvent(conn, emailID, eventID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching attending friends"})
			return
		}

		c.JSON(http.StatusOK, friends)
	}
}
