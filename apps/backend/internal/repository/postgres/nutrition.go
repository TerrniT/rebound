package postgres

import (
	"context"
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

type NutritionRepository struct {
	db *postgres.Postgres
}

func NewNutritionRepository(db *postgres.Postgres) *NutritionRepository {
	return &NutritionRepository{db: db}
}

func (r *NutritionRepository) Create(nutrition *entity.Nutrition) error {
	query := `
		INSERT INTO nutrition (calories, proteins, carbs, fats, serving_size, unit, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	now := time.Now()
	nutrition.CreatedAt = now
	nutrition.UpdatedAt = now

	return r.db.Pool.QueryRow(
		context.Background(),
		query,
		nutrition.Calories,
		nutrition.Proteins,
		nutrition.Carbs,
		nutrition.Fats,
		nutrition.ServingSize,
		nutrition.Unit,
		nutrition.CreatedAt,
		nutrition.UpdatedAt,
	).Scan(&nutrition.ID)
}

func (r *NutritionRepository) GetByID(id int64) (*entity.Nutrition, error) {
	query := `
		SELECT id, calories, proteins, carbs, fats, serving_size, unit, created_at, updated_at
		FROM nutrition
		WHERE id = $1`

	nutrition := &entity.Nutrition{}
	err := r.db.Pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&nutrition.ID,
		&nutrition.Calories,
		&nutrition.Proteins,
		&nutrition.Carbs,
		&nutrition.Fats,
		&nutrition.ServingSize,
		&nutrition.Unit,
		&nutrition.CreatedAt,
		&nutrition.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return nutrition, nil
}

func (r *NutritionRepository) Update(nutrition *entity.Nutrition) error {
	query := `
		UPDATE nutrition
		SET calories = $1, proteins = $2, carbs = $3, fats = $4, serving_size = $5, unit = $6, updated_at = $7
		WHERE id = $8`

	nutrition.UpdatedAt = time.Now()
	_, err := r.db.Pool.Exec(
		context.Background(),
		query,
		nutrition.Calories,
		nutrition.Proteins,
		nutrition.Carbs,
		nutrition.Fats,
		nutrition.ServingSize,
		nutrition.Unit,
		nutrition.UpdatedAt,
		nutrition.ID,
	)
	return err
}

func (r *NutritionRepository) Delete(id int64) error {
	query := `DELETE FROM nutrition WHERE id = $1`
	_, err := r.db.Pool.Exec(context.Background(), query, id)
	return err
}

func (r *NutritionRepository) List(offset, limit int) ([]*entity.Nutrition, error) {
	query := `
		SELECT id, calories, proteins, carbs, fats, serving_size, unit, created_at, updated_at
		FROM nutrition
		ORDER BY id
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nutritions []*entity.Nutrition
	for rows.Next() {
		nutrition := &entity.Nutrition{}
		err := rows.Scan(
			&nutrition.ID,
			&nutrition.Calories,
			&nutrition.Proteins,
			&nutrition.Carbs,
			&nutrition.Fats,
			&nutrition.ServingSize,
			&nutrition.Unit,
			&nutrition.CreatedAt,
			&nutrition.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		nutritions = append(nutritions, nutrition)
	}

	return nutritions, nil
}
