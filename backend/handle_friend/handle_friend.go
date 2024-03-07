package handle_friend

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"

	"github.com/gin-gonic/gin"
)

type User = models.User

// SearchFriends searches for friends based on username and/or name.
// It returns a slice of User structs matching the search criteria.
func SearchFriends(conn *database.DBConnection, searchTerm string) ([]User, error) {
	searchTerm = "%" + strings.ToLower(searchTerm) + "%"
	query := `
	    SELECT email_id, name
	    FROM users
	    WHERE LOWER(name) LIKE $1 OR LOWER(email_id) LIKE $1;
	`
	// query := `
	// SELECT * FROM "users";
	// `
	rows, err := conn.ExecuteQuery(query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundUsers []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.EmailID, &user.Name); err != nil {
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
		searchTerm := c.Query("searchTerm")
		foundUsers, err := SearchFriends(conn, searchTerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching for friends"})
			return
		}
		var friendCards []gin.H
		for _, user := range foundUsers {
			friendCard := gin.H{
				"name":  user.Name,
				"image": "./UCSD-logo.png",
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
