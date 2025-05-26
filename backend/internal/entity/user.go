package entity

import "time"

// User represents a user in the system
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is never exposed in JSON
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository interface defines the methods for user data persistence
type UserRepository interface {
	Create(user *User) error
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id int64) error
}

// UserUseCase interface defines the business logic for users
type UserUseCase interface {
	Register(email, password, name string) (*User, error)
	Login(email, password string) (string, error) // Returns JWT token
	GetUser(id int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
	ValidateToken(token string) (*User, error)
}
