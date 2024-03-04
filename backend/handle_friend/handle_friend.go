package handle_friend

import (
	"log"
	"strings"
	"take-a-break/web-service/database"
)

type User struct {
	Email_id string
	Name     string
}

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
		if err := rows.Scan(&user.Email_id, &user.Name); err != nil {
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

func AddFriend(conn *database.DBConnection, emailID1 string, emailID2 string) error {
	query := `
        INSERT INTO friends (email_id1, email_id2)
        VALUES ($1, $2)
    `

	_, err := conn.ExecuteQuery(query, emailID1, emailID2)
	if err != nil {
		log.Println("Error adding friend:", err)
		return err
	}

	log.Printf("Friendship between '%s' and '%s' added successfully", emailID1, emailID2)
	return nil
}
