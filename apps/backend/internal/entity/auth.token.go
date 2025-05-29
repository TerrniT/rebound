package entity

import "time"

// AuthTokenType represents the type of authentication token
type AuthTokenType string

const (
	AuthTokenTypeAccess  AuthTokenType = "access"
	AuthTokenTypeRefresh AuthTokenType = "refresh"
	AuthTokenTypeReset   AuthTokenType = "reset"
)

// AuthToken represents an authentication token
type AuthToken struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	Type      AuthTokenType `json:"type"`
	TokenHash string        `json:"-"`
	ExpiresAt time.Time     `json:"expires_at"`
	IssuedAt  time.Time     `json:"issued_at"`
	IsRevoked bool          `json:"is_revoked"`
}
