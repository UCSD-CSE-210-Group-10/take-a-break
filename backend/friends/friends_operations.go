package friends

import (
	"fmt"
	"log"
	"strings"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
)

// SearchFriends searches for friends based on username and/or name.
// It returns a slice of User structs matching the search criteria.
func SearchFriends(conn *database.DBConnection, searchTerm string, emailID string) ([]UserRequest, error) {
	searchTerm = "%" + strings.ToLower(searchTerm) + "%"
	query := `
	SELECT 
        u.email_id,
        u.name,
		u.avatar,
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
		if err := rows.Scan(&user.EmailID, &user.Name, &user.Avatar, &user.SentRequest); err != nil {
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

func MakeFriends(conn *database.DBConnection, user1_email, user2_email string) error {
	query := `
		INSERT INTO friends (email_id1, email_id2)
		VALUES ($1, $2), ($3, $4)
	`
	// make a bidirectional connection
	_, err := conn.ExecuteQuery(query, user1_email, user2_email, user2_email, user1_email)
	if err != nil {
		return err
	}

	return nil
}

func FetchFriends(conn *database.DBConnection, email_id string) ([]models.User, error) {
	query := `
		SELECT u.email_id, u.name, u.role, u.avatar
		FROM users u
		INNER JOIN friends f
		ON u.email_id = f.email_id2
		WHERE f.email_id1 = $1
		ORDER BY u.name
	`

	rows, err := conn.ExecuteQuery(query, email_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []models.User
	for rows.Next() {
		var friend models.User
		err := rows.Scan(
			&friend.EmailID,
			&friend.Name,
			&friend.Role,
			&friend.Avatar,
		)
		if err != nil {
			return nil, err
		}

		friends = append(friends, friend)
	}

	return friends, nil
}
