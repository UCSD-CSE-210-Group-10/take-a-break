package events

import (
	"errors"
	"net/http"
	"strconv"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Event = models.Event

var events = []Event{}

func GetEvents(c *gin.Context, conn *database.DBConnection) {
	events, err := FetchAllEvents(conn)
	if err != nil {
		c.Error(err)
		utils.HandleInternalServerError(c, "Failed to fetch events from database", err)
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

// PostEvent handles POST requests to create a new event
func PostEvent(c *gin.Context, conn *database.DBConnection) {
	// Parse form data including files
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
		utils.HandleBadRequest(c, "Failed to parse form data", err)
		return
	}

	// Get the uploaded file from the form
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		utils.HandleBadRequest(c, "No image uploaded", err)
		return
	}
	defer file.Close()

	uniqueFilename := strconv.FormatInt(time.Now().UnixNano(), 10)
	filename, err := utils.SaveUploadedFile(c, file, fileHeader, uniqueFilename)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to save image", err)
		return
	}

	formData := map[string]string{
		"title":       c.Request.FormValue("title"),
		"venue":       c.Request.FormValue("venue"),
		"date":        c.Request.FormValue("date"),
		"time":        c.Request.FormValue("time"),
		"description": c.Request.FormValue("description"),
		"tags":        c.Request.FormValue("tags"),
		"filename":    filename,
		"host":        c.Request.FormValue("host"),
		"contact":     c.Request.FormValue("contact"),
	}

	newEvent, err := InsertEventIntoDatabase(conn, formData)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to insert event into database", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newEvent)
}

func GetEventByID(c *gin.Context, conn *database.DBConnection) {
	id := c.Param("id")

	newEvent, err := FetchEventByID(conn, id)
	if err != nil {
		c.Error(err)
		if errors.Is(err, errors.New("event not found")) {
			utils.HandleNotFound(c, "event not found")
		} else {
			utils.HandleInternalServerError(c, "Failed to fetch event from database", err)
		}
		return
	}

	c.IndentedJSON(http.StatusOK, newEvent)
}
