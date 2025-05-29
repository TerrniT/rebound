package entity

import "time"

// UserRole represents a user's role assignment
type UserRole struct {
	UserID     string    `json:"user_id"`
	RoleID     int       `json:"role_id"`
	AssignedAt time.Time `json:"assigned_at"`
}
