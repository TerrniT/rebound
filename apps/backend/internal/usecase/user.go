package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/repository"
)

type UserConfig struct {
	DefaultPageSize int
	MaxPageSize     int
	MinPasswordLen  int
}

type UserUseCase struct {
	repo   repository.UserRepository
	config UserConfig
}

// NewUserUseCase creates a new instance of UserUseCase
func NewUserUseCase(r repository.UserRepository, config UserConfig) *UserUseCase {
	return &UserUseCase{
		repo:   r,
		config: config,
	}
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(ctx context.Context, user *entity.User, password string) (*entity.User, error) {
	// Validate password
	if len(password) < uc.config.MinPasswordLen {
		return nil, ErrInvalidPassword
	}

	// Check if username is already taken
	existingUser, err := uc.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUsernameTaken
	}

	// Check if email is already taken
	existingUser, err = uc.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrEmailTaken
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set user fields
	user.ID = uuid.New().String()
	user.PasswordHash = string(hashedPassword)
	user.IsActive = true
	user.IsEmailVerified = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, user)
}

// GetUser retrieves a user by ID
func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return uc.repo.GetByID(ctx, id)
}

// GetUserByEmail retrieves a user by email
func (uc *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return uc.repo.GetByEmail(ctx, email)
}

// GetUserByUsername retrieves a user by username
func (uc *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return uc.repo.GetByUsername(ctx, username)
}

// UpdateUser updates an existing user
func (uc *UserUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	// Check if user exists
	existingUser, err := uc.repo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return ErrUserNotFound
	}

	// If username is being changed, check if it's already taken
	if user.Username != existingUser.Username {
		existingUser, err = uc.repo.GetByUsername(ctx, user.Username)
		if err != nil {
			return err
		}
		if existingUser != nil {
			return ErrUsernameTaken
		}
	}

	// If email is being changed, check if it's already taken
	if user.Email != existingUser.Email {
		existingUser, err = uc.repo.GetByEmail(ctx, user.Email)
		if err != nil {
			return err
		}
		if existingUser != nil {
			return ErrEmailTaken
		}
	}

	user.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, user)
}

// DeleteUser deletes a user
func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	// Check if user exists
	existingUser, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return ErrUserNotFound
	}

	return uc.repo.Delete(ctx, id)
}

// ListUsers returns a paginated list of users
func (uc *UserUseCase) ListUsers(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.User, int64, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = uc.config.DefaultPageSize
	}
	if pageSize > uc.config.MaxPageSize {
		pageSize = uc.config.MaxPageSize
	}

	// Get total count
	total, err := uc.repo.Count(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	// Get users
	users, err := uc.repo.List(ctx, filters, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdatePassword updates a user's password
func (uc *UserUseCase) UpdatePassword(ctx context.Context, id string, currentPassword, newPassword string) error {
	// Validate new password
	if len(newPassword) < uc.config.MinPasswordLen {
		return ErrInvalidPassword
	}

	// Get user
	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	// Verify current password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(currentPassword))
	if err != nil {
		return ErrInvalidPassword
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password
	user.PasswordHash = string(hashedPassword)
	user.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, user)
}

// UpdateLastLogin updates the last login timestamp for a user
func (uc *UserUseCase) UpdateLastLogin(ctx context.Context, id string) error {
	return uc.repo.UpdateLastLogin(ctx, id)
}

// UpdateEmailVerification updates the email verification status for a user
func (uc *UserUseCase) UpdateEmailVerification(ctx context.Context, id string, isVerified bool) error {
	return uc.repo.UpdateEmailVerification(ctx, id, isVerified)
}

// VerifyPassword verifies a user's password
func (uc *UserUseCase) VerifyPassword(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
