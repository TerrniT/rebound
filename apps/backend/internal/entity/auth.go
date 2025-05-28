package entity

import "time"

// Token represents an authentication token
type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// AuthRepository interface defines the methods for authentication data persistence
type AuthRepository interface {
	CreateToken(userID int64, token string, expiresAt time.Time) error
	GetToken(token string) (*Token, error)
	DeleteToken(token string) error
}

// AuthUseCase interface defines the business logic for authentication
type AuthUseCase interface {
	GenerateToken(user *User) (*Token, error)
	ValidateToken(token string) (*User, error)
	RevokeToken(token string) error
}
