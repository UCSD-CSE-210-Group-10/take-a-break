package users

import (
	"errors"
	"take-a-break/web-service/database"
)

func InsertUserIntoDatabase(conn *database.DBConnection, user User) (User, error) {
	insertQuery := `
		INSERT INTO users (email_id, name, role)
		VALUES ($1, $2, $3) RETURNING
		email_id, name, role
	`

	// Execute the INSERT query using the ExecuteQuery function
	rows, err := conn.ExecuteQuery(
		insertQuery,
		user.EmailID,
		user.Name,
		user.Role,
	)

	if err != nil {
		return User{}, err
	}

	if rows.Next() {
		var newUser User
		err = rows.Scan(
			&newUser.EmailID,
			&newUser.Name,
			&newUser.Role,
		)

		if err != nil {
			return User{}, err
		}

		return newUser, nil
	}

	return User{}, errors.New("internal server error")
}

func FetchUserByEmailID(conn *database.DBConnection, email_id string) (User, error) {
	query := "SELECT * FROM users WHERE email_id = $1"

	rows, err := conn.ExecuteQuery(query, email_id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	// Process query results
	if rows.Next() {
		// Process the row
		var newUser User
		err := rows.Scan(
			&newUser.EmailID,
			&newUser.Name,
			&newUser.Role,
		)
		if err != nil {
			return User{}, err
		}

		return newUser, nil
	}

	return User{}, errors.New("user not found")
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

func FetchFriends(conn *database.DBConnection, email_id string) ([]User, error) {
	query := `
		SELECT u.email_id, u.name, u.role
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

	var friends []User
	for rows.Next() {
		var friend User
		err := rows.Scan(
			&friend.EmailID,
			&friend.Name,
			&friend.Role,
		)
		if err != nil {
			return nil, err
		}

		friends = append(friends, friend)
	}

	return friends, nil
}
