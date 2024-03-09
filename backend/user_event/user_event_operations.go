package user_event

import (
	"errors"
	"take-a-break/web-service/database"
)

func InsertUserEventIntoDatabase(conn *database.DBConnection, userEvent UserEvent) (UserEvent, error) {
	insertQuery := `
	INSERT INTO user_event (email_id, event_id) 
	VALUES ($1, $2) RETURNING
	email_id, event_id
	`
	
	rows, err := conn.ExecuteQuery(
		insertQuery,
		userEvent.EmailID, 
		userEvent.EventID,
	)

	if err != nil {
		return UserEvent{}, err
	}

	if rows.Next() {
		var newUserEvent UserEvent
		err = rows.Scan(
			&newUserEvent.EmailID,
			&newUserEvent.EventID,
		)

		if err != nil {
			return UserEvent{}, err
		}

		return newUserEvent, nil
	}

	return UserEvent{}, errors.New("internal server error")
}

func GetUserEventFromDatabase(conn *database.DBConnection, emailID, eventID string) (UserEvent, error) {
	var userEvent UserEvent

	selectQuery := `
	SELECT email_id, event_id
	FROM user_event
	WHERE email_id = $1 AND event_id = $2
	`

	rows, err := conn.ExecuteQuery(selectQuery, emailID, eventID)
	if err != nil {
		return UserEvent{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&userEvent.EmailID, &userEvent.EventID)
		if err != nil {
			return UserEvent{}, err
		}
		return userEvent, nil
	}

	return UserEvent{}, errors.New("user event not found")
}