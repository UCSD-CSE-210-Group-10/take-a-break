package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"take-a-break/web-service/constants"

	"github.com/gin-gonic/gin"
)

func HandleNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}

func HandleInternalServerError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": message, "details": err.Error()})
}

func HandleBadRequest(c *gin.Context, message string, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message, "details": err.Error()})
}

func SaveUploadedFile(c *gin.Context, file io.Reader, fileHeader *multipart.FileHeader, filename string) (string, error) {
	file_location := filepath.Join(constants.UPLOAD_PATH, filename+filepath.Ext(fileHeader.Filename))
	err := c.SaveUploadedFile(fileHeader, file_location)
	return file_location, err
}
