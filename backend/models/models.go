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
	Avatar  string `json:"avatar"`
}

type UserEvent struct {
	EmailID string `json:"email_id"`
	EventID string `json:"event_id"`
}

type Friend struct {
	EmailID1 string `json:"email_id1"`
	EmailID2 string `json:"email_id2"`
}

type UserRequest struct {
	EmailID     string `json:"email_id"`
	Name        string `json:"name"`
	SentRequest string `json:"has_sent_request"`
	Avatar      string `json:"avatar"`
}

type Config struct {
	ClientID        string
	ClientSecret    string
	AuthURL         string
	TokenURL        string
	RedirectURL     string
	ClientURL       string
	TokenSecret     string
	TokenExpiration int64
	PostURL         string
}
