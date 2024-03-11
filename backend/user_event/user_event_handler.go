package user_event

import (
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
