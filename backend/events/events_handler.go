package events

import (
	"errors"
	"log"
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

func GetEvents(c *gin.Context) {
	conn, err := database.NewDBConnection()
	if err != nil {
		c.Error(err)
		utils.HandleInternalServerError(c, "Failed to connect to the database", err)
		return
	}
	defer conn.Close()

	events, err := fetchAllEvents(conn)
	if err != nil {
		c.Error(err)
		utils.HandleInternalServerError(c, "Failed to fetch events from database", err)
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

// PostEvent handles POST requests to create a new event
func PostEvent(c *gin.Context) {
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

	// Open a database connection
	conn, err := database.NewDBConnection()
	if err != nil {
		log.Fatal(err)
		utils.HandleInternalServerError(c, "Failed to connect to the database", err)
		return
	}
	defer conn.Close()

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

	newEvent, err := insertEventIntoDatabase(conn, formData)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to insert event into database", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newEvent)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("id")

	conn, err := database.NewDBConnection()
	if err != nil {
		c.Error(err)
		utils.HandleInternalServerError(c, "Failed to connect to the database", err)
		return
	}
	defer conn.Close()

	newEvent, err := fetchEventByID(conn, id)
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
