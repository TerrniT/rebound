package postgres

import (
	"context"
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

type UserRepository struct {
	db *postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	query := `
		INSERT INTO users (email, password, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.Pool.QueryRow(
		context.Background(),
		query,
		user.Email,
		user.Password,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}

func (r *UserRepository) GetByID(id int64) (*entity.User, error) {
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE id = $1`

	user := &entity.User{}
	err := r.db.Pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE email = $1`

	user := &entity.User{}
	err := r.db.Pool.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	query := `
		UPDATE users
		SET email = $1, password = $2, name = $3, updated_at = $4
		WHERE id = $5`

	user.UpdatedAt = time.Now()
	_, err := r.db.Pool.Exec(
		context.Background(),
		query,
		user.Email,
		user.Password,
		user.Name,
		user.UpdatedAt,
		user.ID,
	)
	return err
}

func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Pool.Exec(context.Background(), query, id)
	return err
}
