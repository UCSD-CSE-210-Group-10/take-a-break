package friend_request

import (
	"take-a-break/web-service/database"
	"take-a-break/web-service/friends"
	"take-a-break/web-service/models"
)

func SendFriendRequest(conn *database.DBConnection, user1_email, user2_email string) error {
	query := `
		INSERT INTO friend_requests (sender, reciever)
		VALUES ($1, $2)
	`
	rows, err := conn.ExecuteQuery(query, user1_email, user2_email)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func AcceptFriendRequest(conn *database.DBConnection, user1_email, user2_email string) error {
	// Delete the friend request from the table
	query := `
		DELETE FROM friend_requests
		WHERE (sender = $1 AND reciever = $2)
	`
	rows, err := conn.ExecuteQuery(query, user1_email, user2_email)
	if err != nil {
		return err
	}
	defer rows.Close()

	_, err = friends.MakeFriends(conn, user1_email, user2_email)
	if err != nil {
		return err
	}

	return nil
}

func IgnoreFriendRequest(conn *database.DBConnection, user1_email, user2_email string) error {
	query := `
		UPDATE friend_requests
		SET ignored = true
		WHERE (sender = $1 AND reciever = $2)
	`
	rows, err := conn.ExecuteQuery(query, user1_email, user2_email)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func FetchFriendRequest(conn *database.DBConnection, email_id string) ([]models.User, error) {
	query := `
		SELECT u.email_id, u.name, u.role, u.avatar
		FROM users u
		INNER JOIN friend_requests fr
		ON u.email_id = fr.sender
		WHERE fr.reciever = $1 AND fr.ignored = false
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
