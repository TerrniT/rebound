package postgres

import (
	"context"
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

type MealRepository struct {
	db *postgres.Postgres
}

func NewMealRepository(db *postgres.Postgres) *MealRepository {
	return &MealRepository{db: db}
}

func (r *MealRepository) Create(meal *entity.Meal) error {
	query := `
		INSERT INTO meals (user_id, name, type, nutrition_id, date, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	now := time.Now()
	meal.CreatedAt = now
	meal.UpdatedAt = now

	return r.db.Pool.QueryRow(
		context.Background(),
		query,
		meal.UserID,
		meal.Name,
		meal.Type,
		meal.Nutrition.ID,
		meal.Date,
		meal.Description,
		meal.CreatedAt,
		meal.UpdatedAt,
	).Scan(&meal.ID)
}

func (r *MealRepository) GetByID(id int64) (*entity.Meal, error) {
	query := `
		SELECT m.id, m.user_id, m.name, m.type, m.date, m.description, m.created_at, m.updated_at,
			   n.id, n.calories, n.proteins, n.carbs, n.fats, n.serving_size, n.unit, n.created_at, n.updated_at
		FROM meals m
		JOIN nutrition n ON m.nutrition_id = n.id
		WHERE m.id = $1`

	meal := &entity.Meal{}
	err := r.db.Pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&meal.ID,
		&meal.UserID,
		&meal.Name,
		&meal.Type,
		&meal.Date,
		&meal.Description,
		&meal.CreatedAt,
		&meal.UpdatedAt,
		&meal.Nutrition.ID,
		&meal.Nutrition.Calories,
		&meal.Nutrition.Proteins,
		&meal.Nutrition.Carbs,
		&meal.Nutrition.Fats,
		&meal.Nutrition.ServingSize,
		&meal.Nutrition.Unit,
		&meal.Nutrition.CreatedAt,
		&meal.Nutrition.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (r *MealRepository) Update(meal *entity.Meal) error {
	query := `
		UPDATE meals
		SET name = $1, type = $2, nutrition_id = $3, date = $4, description = $5, updated_at = $6
		WHERE id = $7`

	meal.UpdatedAt = time.Now()
	_, err := r.db.Pool.Exec(
		context.Background(),
		query,
		meal.Name,
		meal.Type,
		meal.Nutrition.ID,
		meal.Date,
		meal.Description,
		meal.UpdatedAt,
		meal.ID,
	)
	return err
}

func (r *MealRepository) Delete(id int64) error {
	query := `DELETE FROM meals WHERE id = $1`
	_, err := r.db.Pool.Exec(context.Background(), query, id)
	return err
}

func (r *MealRepository) List(userID int64, startDate, endDate time.Time) ([]*entity.Meal, error) {
	query := `
		SELECT m.id, m.user_id, m.name, m.type, m.date, m.description, m.created_at, m.updated_at,
			   n.id, n.calories, n.proteins, n.carbs, n.fats, n.serving_size, n.unit, n.created_at, n.updated_at
		FROM meals m
		JOIN nutrition n ON m.nutrition_id = n.id
		WHERE m.user_id = $1 AND m.date BETWEEN $2 AND $3
		ORDER BY m.date DESC`

	rows, err := r.db.Pool.Query(context.Background(), query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []*entity.Meal
	for rows.Next() {
		meal := &entity.Meal{}
		err := rows.Scan(
			&meal.ID,
			&meal.UserID,
			&meal.Name,
			&meal.Type,
			&meal.Date,
			&meal.Description,
			&meal.CreatedAt,
			&meal.UpdatedAt,
			&meal.Nutrition.ID,
			&meal.Nutrition.Calories,
			&meal.Nutrition.Proteins,
			&meal.Nutrition.Carbs,
			&meal.Nutrition.Fats,
			&meal.Nutrition.ServingSize,
			&meal.Nutrition.Unit,
			&meal.Nutrition.CreatedAt,
			&meal.Nutrition.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		meals = append(meals, meal)
	}

	return meals, nil
}

func (r *MealRepository) GetByDate(userID int64, date time.Time) ([]*entity.Meal, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	return r.List(userID, startOfDay, endOfDay)
}
