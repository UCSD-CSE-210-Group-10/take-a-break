package friends

import (
	"net/http"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"

	"github.com/gin-gonic/gin"
)

type User = models.User
type UserRequest = models.UserRequest

func SearchFriendsHandler(conn *database.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Param("token")
		claims := auth.ReturnJWTToken(token)
		emailID := claims["email"].(string)

		searchTerm := c.Query("searchTerm")
		foundUsers, err := SearchFriends(conn, searchTerm, emailID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching for friends"})
			return
		}

		var friendCards []gin.H
		for _, user := range foundUsers {
			friendCard := gin.H{
				"name":             user.Name,
				"email":            user.EmailID,
				"has_sent_request": user.SentRequest,
				"avatar":           user.Avatar,
			}
			friendCards = append(friendCards, friendCard)
		}

		c.JSON(http.StatusOK, friendCards)
	}
}

// Not Using this currently
// func DeleteFriendHandler(conn *database.DBConnection) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		emailID1 := c.PostForm("email_id1")
// 		emailID2 := c.PostForm("email_id2")
// 		err := DeleteFriend(conn, emailID1, emailID2)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting friend"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Friendship between '%s' and '%s' deleted successfully", emailID1, emailID2)})
// 	}
// }

// API not used anywhere
// func PostFriends(c *gin.Context, conn *database.DBConnection) {
// 	var friends struct {
// 		EmailID1 string `json:"email_id_1"`
// 		EmailID2 string `json:"email_id_2"`
// 	}
// 	if err := c.ShouldBindJSON(&friends); err != nil {
// 		utils.HandleBadRequest(c, "Failed to parse the request body", err)
// 		return
// 	}
// 	err := MakeFriends(conn, friends.EmailID1, friends.EmailID2)
// 	if err != nil {
// 		utils.HandleInternalServerError(c, "Failed to make friends", err)
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "Friends added successfully",
// 	})
// }

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
