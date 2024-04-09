package models

// User represents a user in the chat application.
type User struct {
	ID       uint   `json:"id" db:"id"` // Tag "db" is used by sqlx for mapping
	Username string `json:"username" db:"username"`
	Password string `json:"password,omitempty" db:"password"` // omitempty to avoid sending password back in responses
}

// LoginRequest represents the required fields for a user login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest represents the required fields for registering a new user.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
