package user_event

import (
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/utils"

	"github.com/gin-gonic/gin"
)

type Event = models.Event
type UserEvent = models.UserEvent

// PostUserEvent handles the POST request to create a new user event (RSVP)
func PostUserEvent(c *gin.Context, conn *database.DBConnection) {
	var userEvent UserEvent

	if err := c.ShouldBindJSON(&userEvent); err != nil {
		utils.HandleBadRequest(c, "Failed to parse the request body", err)
		return
	}

	userEvent, err := InsertUserEventIntoDatabase(conn, userEvent)

	if err != nil {
		utils.HandleInternalServerError(c, "Failed to insert the user_event into the database", err)
		return
	}

	c.JSON(201, userEvent)
}

// GetUserEvent handles the GET request to retrieve a user event by email ID and event ID
func GetUserEvent(c *gin.Context, conn *database.DBConnection) (UserEvent) {
	emailID := c.Param("email_id")
	eventID := c.Param("event_id")

	userEvent, err := GetUserEventFromDatabase(conn, emailID, eventID)

	if err != nil {
		utils.HandleInternalServerError(c, "Failed to retrieve the user_event from the database", err)
		return UserEvent{}
	}

	c.JSON(200, userEvent)

	return userEvent
}