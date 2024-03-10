package users

import (
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
	emailID := c.Param("email_id")
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
	emailID := c.Param("email_id")
	friends, err := FetchFriends(conn, emailID)

	if err != nil {
		return []User{}, err
	}
	c.JSON(200, friends)
	return friends, nil
}
