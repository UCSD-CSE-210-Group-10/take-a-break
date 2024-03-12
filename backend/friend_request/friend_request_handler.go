package friend_request

import (
	"fmt"
	"net/http"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/utils"

	"github.com/gin-gonic/gin"
)

func PostFriendRequest(c *gin.Context, conn *database.DBConnection) {

	token := c.Param("token")

	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return
	}

	claims := auth.ReturnJWTToken(token)

	senderEmailID := claims["email"].(string)
	var friends struct {
		RecieverEmailID string `json:"email_id"`
	}
	if err := c.ShouldBindJSON(&friends); err != nil {
		utils.HandleBadRequest(c, "Failed to parse the request body", err)
		return
	}

	fmt.Println(senderEmailID)
	fmt.Println(friends.RecieverEmailID)

	err := SendFriendRequest(conn, senderEmailID, friends.RecieverEmailID)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to Send Friend Request", err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Friend Request Sent successfully",
	})
}

func GetFriendRequests(c *gin.Context, conn *database.DBConnection) ([]models.User, error) {

	token := c.Param("token")
	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return []models.User{}, nil
	}

	claims := auth.ReturnJWTToken(token)

	recieverEmailID := claims["email"].(string)

	requests, err := FetchFriendRequest(conn, recieverEmailID)
	if err != nil {
		return []models.User{}, err
	}
	c.JSON(200, requests)
	return requests, nil
}

func PostAcceptFriendRequest(c *gin.Context, conn *database.DBConnection) {

	token := c.Param("token")
	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return
	}

	claims := auth.ReturnJWTToken(token)

	confirmerEmailID := claims["email"].(string)

	var request struct {
		SenderEmailID string `json:"email_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.HandleBadRequest(c, "Failed to parse the request body", err)
		return
	}
	err := AcceptFriendRequest(conn, request.SenderEmailID, confirmerEmailID)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to Accept Friend Request", err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Friend Request Accepted successfully",
	})
}

func PostIgnoreFriendRequest(c *gin.Context, conn *database.DBConnection) {

	token := c.Param("token")
	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return
	}

	claims := auth.ReturnJWTToken(token)

	ignorerEmailID := claims["email"].(string)

	var request struct {
		SenderEmailID string `json:"email_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.HandleBadRequest(c, "Failed to parse the request body", err)
		return
	}
	err := IgnoreFriendRequest(conn, request.SenderEmailID, ignorerEmailID)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to Ignore Friend Request", err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Friend Request Ignored successfully",
	})
}
