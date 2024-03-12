package users

import (
	"fmt"
	"net/http"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/utils"

	"github.com/gin-gonic/gin"
)

type User = models.User

func PostUser(c *gin.Context, conn *database.DBConnection) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleBadRequest(c, "Failed to parse the request body", err)
		return
	}
	user, err := InsertUserIntoDatabase(conn, user)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to insert the user into the database", err)
		return
	}
	c.JSON(201, user)
}

func GetUserByEmailID(c *gin.Context, conn *database.DBConnection) (User, error) {

	token := c.Param("token")
	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return User{}, nil
	}

	claims := auth.ReturnJWTToken(token)

	emailID := claims["email"].(string)

	user, err := FetchUserByEmailID(conn, emailID)
	if err != nil {
		return User{}, err
	}
	c.JSON(200, user)
	return user, nil
}

func PostFriends(c *gin.Context, conn *database.DBConnection) {
	var friends struct {
		EmailID1 string `json:"email_id_1"`
		EmailID2 string `json:"email_id_2"`
	}
	if err := c.ShouldBindJSON(&friends); err != nil {
		utils.HandleBadRequest(c, "Failed to parse the request body", err)
		return
	}
	err := MakeFriends(conn, friends.EmailID1, friends.EmailID2)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to make friends", err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Friends added successfully",
	})
}

func GetFriendsByEmailID(c *gin.Context, conn *database.DBConnection) ([]User, error) {

	token := c.Param("token")
	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return []User{}, nil
	}

	claims := auth.ReturnJWTToken(token)

	emailID := claims["email"].(string)
	friends, err := FetchFriends(conn, emailID)
	if err != nil {
		return []User{}, err
	}
	c.JSON(200, friends)
	return friends, nil
}

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

func GetFriendRequests(c *gin.Context, conn *database.DBConnection) ([]User, error) {

	token := c.Param("token")
	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return []User{}, nil
	}

	claims := auth.ReturnJWTToken(token)

	recieverEmailID := claims["email"].(string)

	requests, err := FetchFriendRequest(conn, recieverEmailID)
	if err != nil {
		return []User{}, err
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
