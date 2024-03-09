package models

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Venue       string `json:"venue"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	ImagePath   string `json:"imagepath"`
	Host        string `json:"host"`
	Contact     string `json:"contact"`
}

type User struct {
	EmailID string `json:"email_id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
}

type UserEvent struct {
	EmailID string `json:"email_id"`
	EventID string `json:"event_id"`
}