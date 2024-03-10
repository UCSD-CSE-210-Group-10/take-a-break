package handle_friend

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"

	"github.com/gin-gonic/gin"
)

type User = models.User
type UserRequest = models.UserRequest

// SearchFriends searches for friends based on username and/or name.
// It returns a slice of User structs matching the search criteria.
func SearchFriends(conn *database.DBConnection, searchTerm string, emailID string) ([]UserRequest, error) {
	searchTerm = "%" + strings.ToLower(searchTerm) + "%"
	query := `
	SELECT 
        u.email_id,
        u.name,
        CASE 
          WHEN friend.email_id1 IS NOT NULL THEN 1 
          WHEN fr.sender IS NOT NULL THEN 2
          ELSE 0 END AS has_sent_request
    FROM 
        users u
    LEFT JOIN 
        (SELECT * FROM friend_requests
         WHERE sender = $2) fr
    ON u.email_id = fr.reciever
   LEFT JOIN (
      SELECT * FROM friends
      WHERE email_id1= $2
) friend
   ON u.email_id = friend.email_id2
    WHERE 
        (LOWER(u.name) LIKE $1 OR LOWER(u.email_id) LIKE $1 ) AND u.email_id != $2;
	`
	// query := `
	// SELECT * FROM "users";
	// `
	fmt.Println(searchTerm)
	fmt.Println(emailID)
	rows, err := conn.ExecuteQuery(query, searchTerm, emailID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundUsers []UserRequest
	for rows.Next() {
		var user UserRequest
		if err := rows.Scan(&user.EmailID, &user.Name, &user.SentRequest); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		foundUsers = append(foundUsers, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foundUsers, nil
}

func DeleteFriend(conn *database.DBConnection, emailID1 string, emailID2 string) error {
	query := `
        DELETE FROM friends
        WHERE (email_id1 = $1 AND email_id2 = $2) OR (email_id1 = $2 AND email_id2 = $1)
    `

	_, err := conn.ExecuteQuery(query, emailID1, emailID2)
	if err != nil {
		log.Println("Error deleting friend:", err)
		return err
	}

	log.Printf("Friendship between '%s' and '%s' deleted successfully", emailID1, emailID2)
	return nil
}

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
				"image":            "./UCSD-logo.png",
			}
			friendCards = append(friendCards, friendCard)
		}

		c.JSON(http.StatusOK, friendCards)
	}
}

func DeleteFriendHandler(conn *database.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID1 := c.PostForm("email_id1")
		emailID2 := c.PostForm("email_id2")
		err := DeleteFriend(conn, emailID1, emailID2)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting friend"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Friendship between '%s' and '%s' deleted successfully", emailID1, emailID2)})
	}
}
