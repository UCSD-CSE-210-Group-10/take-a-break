package events

import (
	"errors"
	"take-a-break/web-service/database"
)

func InsertEventIntoDatabase(conn *database.DBConnection, formData map[string]string) (Event, error) {
	insertQuery := `
		INSERT INTO events (title, venue, date, time, description, tags, imagepath, host, contact)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING
		id, title, venue, date, time, description, tags, imagepath, host, contact
	`

	// Execute the INSERT query using the ExecuteQuery function
	rows, err := conn.ExecuteQuery(
		insertQuery,
		formData["title"],
		formData["venue"],
		formData["date"],
		formData["time"],
		formData["description"],
		formData["tags"],
		formData["filename"],
		formData["host"],
		formData["contact"],
	)

	if err != nil {
		return Event{}, err
	}

	if rows.Next() {
		var newEvent Event
		err = rows.Scan(
			&newEvent.ID,
			&newEvent.Title,
			&newEvent.Venue,
			&newEvent.Date,
			&newEvent.Time,
			&newEvent.Description,
			&newEvent.Tags,
			&newEvent.ImagePath,
			&newEvent.Host,
			&newEvent.Contact,
		)

		if err != nil {
			return Event{}, err
		}

		return newEvent, nil
	}

	return Event{}, errors.New("internal server error")
}

func FetchEventByID(conn *database.DBConnection, id string) (Event, error) {
	query := "SELECT * FROM events WHERE id = $1"

	rows, err := conn.ExecuteQuery(query, id)
	if err != nil {
		return Event{}, err
	}
	defer rows.Close()

	// Process query results
	if rows.Next() {
		// Process the row
		var newEvent Event
		err := rows.Scan(
			&newEvent.ID,
			&newEvent.Title,
			&newEvent.Venue,
			&newEvent.Date,
			&newEvent.Time,
			&newEvent.Description,
			&newEvent.Tags,
			&newEvent.ImagePath,
			&newEvent.Host,
			&newEvent.Contact,
		)
		if err != nil {
			return Event{}, err
		}

		return newEvent, nil
	}

	return Event{}, errors.New("event not found")
}

func FetchAllEvents(conn *database.DBConnection) ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := conn.ExecuteQuery(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var newEvent Event
		err := rows.Scan(
			&newEvent.ID,
			&newEvent.Title,
			&newEvent.Venue,
			&newEvent.Date,
			&newEvent.Time,
			&newEvent.Description,
			&newEvent.Tags,
			&newEvent.ImagePath,
			&newEvent.Host,
			&newEvent.Contact,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, newEvent)
	}

	return events, nil
}
