package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.User, error)
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
	UpdateLastLogin(ctx context.Context, id string) error
	UpdateEmailVerification(ctx context.Context, id string, isVerified bool) error
}

// userRepository implements UserRepository
type userRepository struct {
	db *postgres.Postgres
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *postgres.Postgres) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user in the database
func (r *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	query, args, err := r.db.Builder.Insert("users").
		Columns("user_id", "username", "email", "password_hash", "first_name", "last_name", "date_of_birth", "gender", "profile_picture_url", "is_active", "is_email_verified", "created_at", "updated_at").
		Values(user.ID, user.Username, user.Email, user.PasswordHash, user.FirstName, user.LastName, user.DateOfBirth, user.Gender, user.ProfilePictureURL, user.IsActive, user.IsEmailVerified, user.CreatedAt, user.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByID retrieves a user by their ID
func (r *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query, args, err := r.db.Builder.Select("user_id", "username", "email", "password_hash", "first_name", "last_name", "date_of_birth", "gender", "profile_picture_url", "is_active", "is_email_verified", "last_login_at", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"user_id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user entity.User
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Gender, &user.ProfilePictureURL, &user.IsActive, &user.IsEmailVerified, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by their email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query, args, err := r.db.Builder.Select("user_id", "username", "email", "password_hash", "first_name", "last_name", "date_of_birth", "gender", "profile_picture_url", "is_active", "is_email_verified", "last_login_at", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user entity.User
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Gender, &user.ProfilePictureURL, &user.IsActive, &user.IsEmailVerified, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername retrieves a user by their username
func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	query, args, err := r.db.Builder.Select("user_id", "username", "email", "password_hash", "first_name", "last_name", "date_of_birth", "gender", "profile_picture_url", "is_active", "is_email_verified", "last_login_at", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"username": username}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user entity.User
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Gender, &user.ProfilePictureURL, &user.IsActive, &user.IsEmailVerified, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates an existing user in the database
func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	query, args, err := r.db.Builder.Update("users").
		Set("username", user.Username).
		Set("email", user.Email).
		Set("password_hash", user.PasswordHash).
		Set("first_name", user.FirstName).
		Set("last_name", user.LastName).
		Set("date_of_birth", user.DateOfBirth).
		Set("gender", user.Gender).
		Set("profile_picture_url", user.ProfilePictureURL).
		Set("is_active", user.IsActive).
		Set("is_email_verified", user.IsEmailVerified).
		Set("updated_at", user.UpdatedAt).
		Where(squirrel.Eq{"user_id": user.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// Delete deletes a user from the database
func (r *userRepository) Delete(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("users").
		Where(squirrel.Eq{"user_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// List returns a paginated list of users matching the filters
func (r *userRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.User, error) {
	query := r.db.Builder.Select("user_id", "username", "email", "password_hash", "first_name", "last_name", "date_of_birth", "gender", "profile_picture_url", "is_active", "is_email_verified", "last_login_at", "created_at", "updated_at").
		From("users")

	// Apply filters
	for key, value := range filters {
		query = query.Where(squirrel.Eq{key: value})
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	query = query.Limit(uint64(pageSize)).Offset(uint64(offset))

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Gender, &user.ProfilePictureURL, &user.IsActive, &user.IsEmailVerified, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// Count returns the total number of users matching the filters
func (r *userRepository) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	query := r.db.Builder.Select("COUNT(*)").
		From("users")

	// Apply filters
	for key, value := range filters {
		query = query.Where(squirrel.Eq{key: value})
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var count int64
	err = r.db.Pool.QueryRow(ctx, sqlQuery, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// UpdateLastLogin updates the last login timestamp for a user
func (r *userRepository) UpdateLastLogin(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Update("users").
		Set("last_login_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Where(squirrel.Eq{"user_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// UpdateEmailVerification updates the email verification status for a user
func (r *userRepository) UpdateEmailVerification(ctx context.Context, id string, isVerified bool) error {
	query, args, err := r.db.Builder.Update("users").
		Set("is_email_verified", isVerified).
		Where(squirrel.Eq{"user_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}
