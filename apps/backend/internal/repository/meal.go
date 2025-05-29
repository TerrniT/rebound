package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

// MealRepository defines the interface for meal-related database operations
type MealRepository interface {
	Create(ctx context.Context, meal *entity.UserMeal) (*entity.UserMeal, error)
	GetByID(ctx context.Context, id string) (*entity.UserMeal, error)
	Update(ctx context.Context, meal *entity.UserMeal) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.UserMeal, error)
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.UserMeal, error)
	AddFoodItem(ctx context.Context, foodItem *entity.MealFoodItem) error
	GetFoodItems(ctx context.Context, mealID string) ([]*entity.MealFoodItem, error)
	UpdateFoodItem(ctx context.Context, foodItem *entity.MealFoodItem) error
	DeleteFoodItem(ctx context.Context, foodItemID string) error
}

// mealRepository implements MealRepository
type mealRepository struct {
	db *postgres.Postgres
}

// NewMealRepository creates a new instance of MealRepository
func NewMealRepository(db *postgres.Postgres) MealRepository {
	return &mealRepository{db: db}
}

// Create creates a new meal in the database
func (r *mealRepository) Create(ctx context.Context, meal *entity.UserMeal) (*entity.UserMeal, error) {
	query, args, err := r.db.Builder.Insert("user_meals").
		Columns("meal_id", "user_id", "meal_type", "meal_date", "meal_time", "custom_meal_name", "notes", "total_calories_consumed", "total_protein_consumed", "total_fat_consumed", "total_carbs_consumed", "created_at", "updated_at").
		Values(meal.ID, meal.UserID, meal.MealType, meal.MealDate, meal.MealTime, meal.CustomMealName, meal.Notes, meal.TotalCaloriesConsumed, meal.TotalProteinConsumed, meal.TotalFatConsumed, meal.TotalCarbsConsumed, meal.CreatedAt, meal.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return meal, nil
}

// GetByID retrieves a meal by its ID
func (r *mealRepository) GetByID(ctx context.Context, id string) (*entity.UserMeal, error) {
	query, args, err := r.db.Builder.Select("meal_id", "user_id", "meal_type", "meal_date", "meal_time", "custom_meal_name", "notes", "total_calories_consumed", "total_protein_consumed", "total_fat_consumed", "total_carbs_consumed", "created_at", "updated_at").
		From("user_meals").
		Where(squirrel.Eq{"meal_id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var meal entity.UserMeal
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&meal.ID, &meal.UserID, &meal.MealType, &meal.MealDate, &meal.MealTime, &meal.CustomMealName, &meal.Notes, &meal.TotalCaloriesConsumed, &meal.TotalProteinConsumed, &meal.TotalFatConsumed, &meal.TotalCarbsConsumed, &meal.CreatedAt, &meal.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &meal, nil
}

// Update updates an existing meal in the database
func (r *mealRepository) Update(ctx context.Context, meal *entity.UserMeal) error {
	query, args, err := r.db.Builder.Update("user_meals").
		Set("user_id", meal.UserID).
		Set("meal_type", meal.MealType).
		Set("meal_date", meal.MealDate).
		Set("meal_time", meal.MealTime).
		Set("custom_meal_name", meal.CustomMealName).
		Set("notes", meal.Notes).
		Set("total_calories_consumed", meal.TotalCaloriesConsumed).
		Set("total_protein_consumed", meal.TotalProteinConsumed).
		Set("total_fat_consumed", meal.TotalFatConsumed).
		Set("total_carbs_consumed", meal.TotalCarbsConsumed).
		Set("updated_at", meal.UpdatedAt).
		Where(squirrel.Eq{"meal_id": meal.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// Delete deletes a meal from the database
func (r *mealRepository) Delete(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("user_meals").
		Where(squirrel.Eq{"meal_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// List returns a paginated list of meals matching the filters
func (r *mealRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.UserMeal, error) {
	query := r.db.Builder.Select("meal_id", "user_id", "meal_type", "meal_date", "meal_time", "custom_meal_name", "notes", "total_calories_consumed", "total_protein_consumed", "total_fat_consumed", "total_carbs_consumed", "created_at", "updated_at").
		From("user_meals")

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

	var meals []*entity.UserMeal
	for rows.Next() {
		var meal entity.UserMeal
		err := rows.Scan(
			&meal.ID, &meal.UserID, &meal.MealType, &meal.MealDate, &meal.MealTime, &meal.CustomMealName, &meal.Notes, &meal.TotalCaloriesConsumed, &meal.TotalProteinConsumed, &meal.TotalFatConsumed, &meal.TotalCarbsConsumed, &meal.CreatedAt, &meal.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		meals = append(meals, &meal)
	}

	return meals, nil
}

// Count returns the total number of meals matching the filters
func (r *mealRepository) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	query := r.db.Builder.Select("COUNT(*)").
		From("user_meals")

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

// GetByUserID retrieves meals for a specific user
func (r *mealRepository) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.UserMeal, error) {
	query, args, err := r.db.Builder.Select("meal_id", "user_id", "meal_type", "meal_date", "meal_time", "custom_meal_name", "notes", "total_calories_consumed", "total_protein_consumed", "total_fat_consumed", "total_carbs_consumed", "created_at", "updated_at").
		From("user_meals").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("meal_date DESC", "meal_time DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var meals []*entity.UserMeal
	for rows.Next() {
		var meal entity.UserMeal
		err := rows.Scan(
			&meal.ID, &meal.UserID, &meal.MealType, &meal.MealDate, &meal.MealTime, &meal.CustomMealName, &meal.Notes, &meal.TotalCaloriesConsumed, &meal.TotalProteinConsumed, &meal.TotalFatConsumed, &meal.TotalCarbsConsumed, &meal.CreatedAt, &meal.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		meals = append(meals, &meal)
	}
	return meals, nil
}

// AddFoodItem adds a new food item to a meal
func (r *mealRepository) AddFoodItem(ctx context.Context, foodItem *entity.MealFoodItem) error {
	query, args, err := r.db.Builder.Insert("meal_food_items").
		Columns("meal_food_item_id", "meal_id", "food_item_id", "quantity_consumed", "serving_unit_consumed", "calories_consumed", "protein_consumed", "fat_consumed", "carbs_consumed", "logged_at").
		Values(foodItem.ID, foodItem.MealID, foodItem.FoodItemID, foodItem.QuantityConsumed, foodItem.ServingUnitConsumed, foodItem.CaloriesConsumed, foodItem.ProteinConsumed, foodItem.FatConsumed, foodItem.CarbsConsumed, foodItem.LoggedAt).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// GetFoodItems retrieves all food items for a meal
func (r *mealRepository) GetFoodItems(ctx context.Context, mealID string) ([]*entity.MealFoodItem, error) {
	query, args, err := r.db.Builder.Select("meal_food_item_id", "meal_id", "food_item_id", "quantity_consumed", "serving_unit_consumed", "calories_consumed", "protein_consumed", "fat_consumed", "carbs_consumed", "logged_at").
		From("meal_food_items").
		Where(squirrel.Eq{"meal_id": mealID}).
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var foodItems []*entity.MealFoodItem
	for rows.Next() {
		var foodItem entity.MealFoodItem
		err := rows.Scan(
			&foodItem.ID, &foodItem.MealID, &foodItem.FoodItemID, &foodItem.QuantityConsumed, &foodItem.ServingUnitConsumed, &foodItem.CaloriesConsumed, &foodItem.ProteinConsumed, &foodItem.FatConsumed, &foodItem.CarbsConsumed, &foodItem.LoggedAt,
		)
		if err != nil {
			return nil, err
		}
		foodItems = append(foodItems, &foodItem)
	}
	return foodItems, nil
}

// UpdateFoodItem updates an existing food item
func (r *mealRepository) UpdateFoodItem(ctx context.Context, foodItem *entity.MealFoodItem) error {
	query, args, err := r.db.Builder.Update("meal_food_items").
		Set("quantity_consumed", foodItem.QuantityConsumed).
		Set("serving_unit_consumed", foodItem.ServingUnitConsumed).
		Set("calories_consumed", foodItem.CaloriesConsumed).
		Set("protein_consumed", foodItem.ProteinConsumed).
		Set("fat_consumed", foodItem.FatConsumed).
		Set("carbs_consumed", foodItem.CarbsConsumed).
		Where(squirrel.Eq{"meal_food_item_id": foodItem.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// DeleteFoodItem deletes a food item
func (r *mealRepository) DeleteFoodItem(ctx context.Context, foodItemID string) error {
	query, args, err := r.db.Builder.Delete("meal_food_items").
		Where(squirrel.Eq{"meal_food_item_id": foodItemID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}
