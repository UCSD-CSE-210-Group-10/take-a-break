package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var events = []event{}

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func postEvent(c *gin.Context) {
	var newEvent event

	// Parse form data including files
	if err := c.Request.ParseMultipartForm(10 << 20); // 10 MB limit
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the uploaded file from the form
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded"})
		return
	}
	defer file.Close()

	eventId := strconv.FormatInt(time.Now().UnixNano(), 10)

	uploadPath := "./event-posters/"
	filename := filepath.Join(uploadPath, eventId+filepath.Ext(fileHeader.Filename))
	err = c.SaveUploadedFile(fileHeader, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	fmt.Println("")

	// if the field does not exist, PostForm returns an empty string
	newEvent.ID = eventId
	newEvent.ImagePath = filename
	newEvent.Title = c.PostForm("title")
	newEvent.Venue = c.PostForm("venue")
	newEvent.Date = c.PostForm("date")
	newEvent.Time = c.PostForm("time")
	newEvent.Description = c.PostForm("description")
	newEvent.Tags = c.PostForm("tags")
	newEvent.Host = c.PostForm("host")
	newEvent.Contact = c.PostForm("contact")

	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func getEventByID(c *gin.Context) {
	id := c.Param("id")

	for _, e := range events {
		if e.ID == id {
			c.IndentedJSON(http.StatusOK, e)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}
