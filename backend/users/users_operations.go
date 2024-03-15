package users

import (
	"errors"
	"take-a-break/web-service/database"
)

func InsertUserIntoDatabase(conn *database.DBConnection, user User) (User, error) {
	insertQuery := `
		INSERT INTO users (email_id, name, role, avatar)
		VALUES ($1, $2, $3, $4) RETURNING
		email_id, name, role, avatar
	`

	// Execute the INSERT query using the ExecuteQuery function
	rows, err := conn.ExecuteQuery(
		insertQuery,
		user.EmailID,
		user.Name,
		user.Role,
		user.Avatar,
	)
	defer rows.Close()

	if err != nil {
		return User{}, err
	}

	if rows.Next() {
		var newUser User
		err = rows.Scan(
			&newUser.EmailID,
			&newUser.Name,
			&newUser.Role,
			&newUser.Avatar,
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
			&newUser.Avatar,
		)
		if err != nil {
			return User{}, err
		}

		return newUser, nil
	}

	return User{}, errors.New("user not found")
}
