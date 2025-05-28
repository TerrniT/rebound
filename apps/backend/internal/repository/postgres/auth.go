package postgres

import (
	"context"
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

type AuthRepository struct {
	db *postgres.Postgres
}

func NewAuthRepository(db *postgres.Postgres) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateToken(userID int64, token string, expiresAt time.Time) error {
	query := `
		INSERT INTO tokens (user_id, token, expires_at, created_at)
		VALUES ($1, $2, $3, $4)`

	_, err := r.db.Pool.Exec(
		context.Background(),
		query,
		userID,
		token,
		expiresAt,
		time.Now(),
	)
	return err
}

func (r *AuthRepository) GetToken(token string) (*entity.Token, error) {
	query := `
		SELECT token, expires_at
		FROM tokens
		WHERE token = $1 AND expires_at > $2`

	tokenEntity := &entity.Token{}
	err := r.db.Pool.QueryRow(
		context.Background(),
		query,
		token,
		time.Now(),
	).Scan(
		&tokenEntity.AccessToken,
		&tokenEntity.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}

	return tokenEntity, nil
}

func (r *AuthRepository) DeleteToken(token string) error {
	query := `DELETE FROM tokens WHERE token = $1`
	_, err := r.db.Pool.Exec(context.Background(), query, token)
	return err
}
