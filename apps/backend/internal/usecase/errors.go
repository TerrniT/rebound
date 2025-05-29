package usecase

import "errors"

var (
	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("user not found")

	// ErrInvalidPassword is returned when a password is invalid
	ErrInvalidPassword = errors.New("invalid password")

	// ErrUsernameTaken is returned when a username is already taken
	ErrUsernameTaken = errors.New("username already taken")

	// ErrEmailTaken is returned when an email is already taken
	ErrEmailTaken = errors.New("email already taken")

	// ErrInvalidInput is returned when input validation fails
	ErrInvalidInput = errors.New("invalid input")

	// ErrUnauthorized is returned when a user is not authorized to perform an action
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden is returned when a user is forbidden from performing an action
	ErrForbidden = errors.New("forbidden")

	// ErrInternal is returned when an internal error occurs
	ErrInternal = errors.New("internal error")
)
