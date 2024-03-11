package user_event

import (
	"errors"
	"net/http"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/utils"

	"github.com/gin-gonic/gin"
)

type Event = models.Event
type UserEvent = models.UserEvent

// PostUserEvent handles the POST request to create a new user event (RSVP)
func PostUserEvent(c *gin.Context, conn *database.DBConnection) {
	token := c.Param("token")
	eventID := c.Param("event_id")

	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
	}

	claims := auth.ReturnJWTToken(token)

	emailID := claims["email"].(string)

	userEvent, err := InsertUserEventIntoDatabase(conn, emailID, eventID)

	if err != nil {
		utils.HandleInternalServerError(c, "Failed to insert the user_event into the database", err)
		return
	}

	c.JSON(201, userEvent)
}

func InsertUserEventIntoDatabase(conn *database.DBConnection, emailID string, eventID string) (UserEvent, error) {
	insertQuery := `
	INSERT INTO user_event (email_id, event_id) 
	VALUES ($1, $2) RETURNING
	email_id, event_id
	`

	rows, err := conn.ExecuteQuery(
		insertQuery,
		emailID,
		eventID,
	)

	if err != nil {
		return UserEvent{}, err
	}

	if rows.Next() {
		var newUserEvent UserEvent
		err = rows.Scan(
			&newUserEvent.EmailID,
			&newUserEvent.EventID,
		)

		if err != nil {
			return UserEvent{}, err
		}

		return newUserEvent, nil
	}

	return UserEvent{}, errors.New("internal server error")
}

// GetUserEvent handles the GET request to retrieve a user event by email ID and event ID
func GetUserEvent(c *gin.Context, conn *database.DBConnection) UserEvent {
	token := c.Param("token")
	eventID := c.Param("event_id")

	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
	}

	claims := auth.ReturnJWTToken(token)

	emailID := claims["email"].(string)

	userEvent, err := GetUserEventFromDatabase(conn, emailID, eventID)

	if err != nil {
		utils.HandleInternalServerError(c, "Failed to retrieve the user_event from the database", err)
		return UserEvent{}
	}

	c.JSON(200, userEvent)

	return userEvent
}

func GetUserEventFromDatabase(conn *database.DBConnection, emailID string, eventID string) (UserEvent, error) {
	var userEvent UserEvent

	selectQuery := `
	SELECT email_id, event_id
	FROM user_event
	WHERE email_id = $1 AND event_id = $2
	`

	rows, err := conn.ExecuteQuery(selectQuery, emailID, eventID)
	if err != nil {
		return UserEvent{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&userEvent.EmailID, &userEvent.EventID)
		if err != nil {
			return UserEvent{}, err
		}
		return userEvent, nil
	}

	return UserEvent{}, errors.New("user event not found")
}
