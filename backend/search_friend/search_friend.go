package search_friend

import (
	"log"
	"strings"
)

// SearchFriends searches for friends based on username and/or name.
// It returns a slice of User structs matching the search criteria.
func (conn *DBConnection) SearchFriends(searchTerm string) ([]User, error) {
	searchTerm = "%" + strings.ToLower(searchTerm) + "%"
	query := `
        SELECT username, name
        FROM users
        WHERE LOWER(username) LIKE $1 OR LOWER(name) LIKE $1
    `

	rows, err := conn.ExecuteQuery(query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foundUsers []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username, &user.Name); err != nil {
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
