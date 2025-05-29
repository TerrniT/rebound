package entity

import "time"

// UserGender represents the gender of a user
type UserGender string

const (
	UserGenderMale   UserGender = "Male"
	UserGenderFemale UserGender = "Female"
)

// User represents a user in the system
type User struct {
	ID                string     `json:"id"`
	Username          string     `json:"username"`
	Email             string     `json:"email"`
	PasswordHash      string     `json:"-"` // Never expose password hash in JSON
	FirstName         string     `json:"first_name"`
	LastName          string     `json:"last_name"`
	DateOfBirth       *time.Time `json:"date_of_birth,omitempty"`
	Gender            UserGender `json:"gender,omitempty"`
	ProfilePictureURL string     `json:"profile_picture_url,omitempty"`
	IsActive          bool       `json:"is_active"`
	IsEmailVerified   bool       `json:"is_email_verified"`
	LastLoginAt       *time.Time `json:"last_login_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
