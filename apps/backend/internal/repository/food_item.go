package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/terrnit/rebound/backend/internal/entity"

	"github.com/terrnit/rebound/backend/pkg/postgres"
)

// FoodItemRepository defines the interface for food item-related database operations
type FoodItemRepository interface {
	Create(ctx context.Context, foodItem *entity.FoodItem) (*entity.FoodItem, error)
	GetByID(ctx context.Context, id string) (*entity.FoodItem, error)
	GetByBarcode(ctx context.Context, barcode string) (*entity.FoodItem, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*entity.FoodItem, error)
	Update(ctx context.Context, foodItem *entity.FoodItem) error
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.FoodItem, error)
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
	CountBySearch(ctx context.Context, query string) (int64, error)
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.FoodItem, error)
}

// foodItemRepository implements FoodItemRepository
type foodItemRepository struct {
	db *postgres.Postgres
}

// NewFoodItemRepository creates a new instance of FoodItemRepository
func NewFoodItemRepository(db *postgres.Postgres) FoodItemRepository {
	return &foodItemRepository{db: db}
}

// Create creates a new food item in the database
func (r *foodItemRepository) Create(ctx context.Context, foodItem *entity.FoodItem) (*entity.FoodItem, error) {
	query, args, err := r.db.Builder.Insert("food_items").
		Columns("id", "name", "brand_name", "barcode_upc", "serving_size_default_qty", "serving_size_default_unit", "calories_per_default_serving", "protein_grams_per_default_serving", "fat_grams_per_default_serving", "carbs_grams_per_default_serving", "fiber_grams_per_default_serving", "sugar_grams_per_default_serving", "saturated_fat_grams_per_default_serving", "trans_fat_grams_per_default_serving", "cholesterol_mg_per_default_serving", "sodium_mg_per_default_serving", "potassium_mg_per_default_serving", "vitamin_a_mcg_per_default_serving", "vitamin_c_mg_per_default_serving", "calcium_mg_per_default_serving", "iron_mg_per_default_serving", "source", "is_verified", "created_by_user_id", "created_at", "updated_at").
		Values(foodItem.ID, foodItem.Name, foodItem.BrandName, foodItem.BarcodeUPC, foodItem.ServingSizeDefaultQty, foodItem.ServingSizeDefaultUnit, foodItem.CaloriesPerDefaultServing, foodItem.ProteinGramsPerDefaultServing, foodItem.FatGramsPerDefaultServing, foodItem.CarbsGramsPerDefaultServing, foodItem.FiberGramsPerDefaultServing, foodItem.SugarGramsPerDefaultServing, foodItem.SaturatedFatGramsPerDefaultServing, foodItem.TransFatGramsPerDefaultServing, foodItem.CholesterolMgPerDefaultServing, foodItem.SodiumMgPerDefaultServing, foodItem.PotassiumMgPerDefaultServing, foodItem.VitaminAMcgPerDefaultServing, foodItem.VitaminCMgPerDefaultServing, foodItem.CalciumMgPerDefaultServing, foodItem.IronMgPerDefaultServing, foodItem.Source, foodItem.IsVerified, foodItem.CreatedByUserID, foodItem.CreatedAt, foodItem.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return foodItem, nil
}

// GetByID retrieves a food item by its ID
func (r *foodItemRepository) GetByID(ctx context.Context, id string) (*entity.FoodItem, error) {
	query, args, err := r.db.Builder.Select("id", "name", "brand_name", "barcode_upc", "serving_size_default_qty", "serving_size_default_unit", "calories_per_default_serving", "protein_grams_per_default_serving", "fat_grams_per_default_serving", "carbs_grams_per_default_serving", "fiber_grams_per_default_serving", "sugar_grams_per_default_serving", "saturated_fat_grams_per_default_serving", "trans_fat_grams_per_default_serving", "cholesterol_mg_per_default_serving", "sodium_mg_per_default_serving", "potassium_mg_per_default_serving", "vitamin_a_mcg_per_default_serving", "vitamin_c_mg_per_default_serving", "calcium_mg_per_default_serving", "iron_mg_per_default_serving", "source", "is_verified", "created_by_user_id", "created_at", "updated_at").
		From("food_items").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var foodItem entity.FoodItem
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&foodItem.ID, &foodItem.Name, &foodItem.BrandName, &foodItem.BarcodeUPC, &foodItem.ServingSizeDefaultQty, &foodItem.ServingSizeDefaultUnit, &foodItem.CaloriesPerDefaultServing, &foodItem.ProteinGramsPerDefaultServing, &foodItem.FatGramsPerDefaultServing, &foodItem.CarbsGramsPerDefaultServing, &foodItem.FiberGramsPerDefaultServing, &foodItem.SugarGramsPerDefaultServing, &foodItem.SaturatedFatGramsPerDefaultServing, &foodItem.TransFatGramsPerDefaultServing, &foodItem.CholesterolMgPerDefaultServing, &foodItem.SodiumMgPerDefaultServing, &foodItem.PotassiumMgPerDefaultServing, &foodItem.VitaminAMcgPerDefaultServing, &foodItem.VitaminCMgPerDefaultServing, &foodItem.CalciumMgPerDefaultServing, &foodItem.IronMgPerDefaultServing, &foodItem.Source, &foodItem.IsVerified, &foodItem.CreatedByUserID, &foodItem.CreatedAt, &foodItem.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &foodItem, nil
}

// GetByBarcode retrieves a food item by its barcode
func (r *foodItemRepository) GetByBarcode(ctx context.Context, barcode string) (*entity.FoodItem, error) {
	query, args, err := r.db.Builder.Select("id", "name", "brand_name", "barcode_upc", "serving_size_default_qty", "serving_size_default_unit", "calories_per_default_serving", "protein_grams_per_default_serving", "fat_grams_per_default_serving", "carbs_grams_per_default_serving", "fiber_grams_per_default_serving", "sugar_grams_per_default_serving", "saturated_fat_grams_per_default_serving", "trans_fat_grams_per_default_serving", "cholesterol_mg_per_default_serving", "sodium_mg_per_default_serving", "potassium_mg_per_default_serving", "vitamin_a_mcg_per_default_serving", "vitamin_c_mg_per_default_serving", "calcium_mg_per_default_serving", "iron_mg_per_default_serving", "source", "is_verified", "created_by_user_id", "created_at", "updated_at").
		From("food_items").
		Where(squirrel.Eq{"barcode_upc": barcode}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var foodItem entity.FoodItem
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&foodItem.ID, &foodItem.Name, &foodItem.BrandName, &foodItem.BarcodeUPC, &foodItem.ServingSizeDefaultQty, &foodItem.ServingSizeDefaultUnit, &foodItem.CaloriesPerDefaultServing, &foodItem.ProteinGramsPerDefaultServing, &foodItem.FatGramsPerDefaultServing, &foodItem.CarbsGramsPerDefaultServing, &foodItem.FiberGramsPerDefaultServing, &foodItem.SugarGramsPerDefaultServing, &foodItem.SaturatedFatGramsPerDefaultServing, &foodItem.TransFatGramsPerDefaultServing, &foodItem.CholesterolMgPerDefaultServing, &foodItem.SodiumMgPerDefaultServing, &foodItem.PotassiumMgPerDefaultServing, &foodItem.VitaminAMcgPerDefaultServing, &foodItem.VitaminCMgPerDefaultServing, &foodItem.CalciumMgPerDefaultServing, &foodItem.IronMgPerDefaultServing, &foodItem.Source, &foodItem.IsVerified, &foodItem.CreatedByUserID, &foodItem.CreatedAt, &foodItem.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &foodItem, nil
}

// Search searches for food items by name or brand
func (r *foodItemRepository) Search(ctx context.Context, query string, limit, offset int) ([]*entity.FoodItem, error) {
	sqlQuery, args, err := r.db.Builder.Select("id", "name", "brand_name", "barcode_upc", "serving_size_default_qty", "serving_size_default_unit", "calories_per_default_serving", "protein_grams_per_default_serving", "fat_grams_per_default_serving", "carbs_grams_per_default_serving", "fiber_grams_per_default_serving", "sugar_grams_per_default_serving", "saturated_fat_grams_per_default_serving", "trans_fat_grams_per_default_serving", "cholesterol_mg_per_default_serving", "sodium_mg_per_default_serving", "potassium_mg_per_default_serving", "vitamin_a_mcg_per_default_serving", "vitamin_c_mg_per_default_serving", "calcium_mg_per_default_serving", "iron_mg_per_default_serving", "source", "is_verified", "created_by_user_id", "created_at", "updated_at").
		From("food_items").
		Where(squirrel.Or{
			squirrel.Like{"name": "%" + query + "%"},
			squirrel.Like{"brand_name": "%" + query + "%"},
		}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var foodItems []*entity.FoodItem
	for rows.Next() {
		var foodItem entity.FoodItem
		err := rows.Scan(
			&foodItem.ID, &foodItem.Name, &foodItem.BrandName, &foodItem.BarcodeUPC, &foodItem.ServingSizeDefaultQty, &foodItem.ServingSizeDefaultUnit, &foodItem.CaloriesPerDefaultServing, &foodItem.ProteinGramsPerDefaultServing, &foodItem.FatGramsPerDefaultServing, &foodItem.CarbsGramsPerDefaultServing, &foodItem.FiberGramsPerDefaultServing, &foodItem.SugarGramsPerDefaultServing, &foodItem.SaturatedFatGramsPerDefaultServing, &foodItem.TransFatGramsPerDefaultServing, &foodItem.CholesterolMgPerDefaultServing, &foodItem.SodiumMgPerDefaultServing, &foodItem.PotassiumMgPerDefaultServing, &foodItem.VitaminAMcgPerDefaultServing, &foodItem.VitaminCMgPerDefaultServing, &foodItem.CalciumMgPerDefaultServing, &foodItem.IronMgPerDefaultServing, &foodItem.Source, &foodItem.IsVerified, &foodItem.CreatedByUserID, &foodItem.CreatedAt, &foodItem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		foodItems = append(foodItems, &foodItem)
	}
	return foodItems, nil
}

// Update updates an existing food item in the database
func (r *foodItemRepository) Update(ctx context.Context, foodItem *entity.FoodItem) error {
	query, args, err := r.db.Builder.Update("food_items").
		Set("name", foodItem.Name).
		Set("brand_name", foodItem.BrandName).
		Set("barcode_upc", foodItem.BarcodeUPC).
		Set("serving_size_default_qty", foodItem.ServingSizeDefaultQty).
		Set("serving_size_default_unit", foodItem.ServingSizeDefaultUnit).
		Set("calories_per_default_serving", foodItem.CaloriesPerDefaultServing).
		Set("protein_grams_per_default_serving", foodItem.ProteinGramsPerDefaultServing).
		Set("fat_grams_per_default_serving", foodItem.FatGramsPerDefaultServing).
		Set("carbs_grams_per_default_serving", foodItem.CarbsGramsPerDefaultServing).
		Set("fiber_grams_per_default_serving", foodItem.FiberGramsPerDefaultServing).
		Set("sugar_grams_per_default_serving", foodItem.SugarGramsPerDefaultServing).
		Set("saturated_fat_grams_per_default_serving", foodItem.SaturatedFatGramsPerDefaultServing).
		Set("trans_fat_grams_per_default_serving", foodItem.TransFatGramsPerDefaultServing).
		Set("cholesterol_mg_per_default_serving", foodItem.CholesterolMgPerDefaultServing).
		Set("sodium_mg_per_default_serving", foodItem.SodiumMgPerDefaultServing).
		Set("potassium_mg_per_default_serving", foodItem.PotassiumMgPerDefaultServing).
		Set("vitamin_a_mcg_per_default_serving", foodItem.VitaminAMcgPerDefaultServing).
		Set("vitamin_c_mg_per_default_serving", foodItem.VitaminCMgPerDefaultServing).
		Set("calcium_mg_per_default_serving", foodItem.CalciumMgPerDefaultServing).
		Set("iron_mg_per_default_serving", foodItem.IronMgPerDefaultServing).
		Set("source", foodItem.Source).
		Set("is_verified", foodItem.IsVerified).
		Set("created_by_user_id", foodItem.CreatedByUserID).
		Set("updated_at", foodItem.UpdatedAt).
		Where(squirrel.Eq{"id": foodItem.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// Delete deletes a food item from the database
func (r *foodItemRepository) Delete(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("food_items").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// GetByUserID retrieves food items created by a specific user
func (r *foodItemRepository) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.FoodItem, error) {
	query, args, err := r.db.Builder.Select("id", "name", "brand_name", "barcode_upc", "serving_size_default_qty", "serving_size_default_unit", "calories_per_default_serving", "protein_grams_per_default_serving", "fat_grams_per_default_serving", "carbs_grams_per_default_serving", "fiber_grams_per_default_serving", "sugar_grams_per_default_serving", "saturated_fat_grams_per_default_serving", "trans_fat_grams_per_default_serving", "cholesterol_mg_per_default_serving", "sodium_mg_per_default_serving", "potassium_mg_per_default_serving", "vitamin_a_mcg_per_default_serving", "vitamin_c_mg_per_default_serving", "calcium_mg_per_default_serving", "iron_mg_per_default_serving", "source", "is_verified", "created_by_user_id", "created_at", "updated_at").
		From("food_items").
		Where(squirrel.Eq{"created_by_user_id": userID}).
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
	var foodItems []*entity.FoodItem
	for rows.Next() {
		var foodItem entity.FoodItem
		err := rows.Scan(
			&foodItem.ID, &foodItem.Name, &foodItem.BrandName, &foodItem.BarcodeUPC, &foodItem.ServingSizeDefaultQty, &foodItem.ServingSizeDefaultUnit, &foodItem.CaloriesPerDefaultServing, &foodItem.ProteinGramsPerDefaultServing, &foodItem.FatGramsPerDefaultServing, &foodItem.CarbsGramsPerDefaultServing, &foodItem.FiberGramsPerDefaultServing, &foodItem.SugarGramsPerDefaultServing, &foodItem.SaturatedFatGramsPerDefaultServing, &foodItem.TransFatGramsPerDefaultServing, &foodItem.CholesterolMgPerDefaultServing, &foodItem.SodiumMgPerDefaultServing, &foodItem.PotassiumMgPerDefaultServing, &foodItem.VitaminAMcgPerDefaultServing, &foodItem.VitaminCMgPerDefaultServing, &foodItem.CalciumMgPerDefaultServing, &foodItem.IronMgPerDefaultServing, &foodItem.Source, &foodItem.IsVerified, &foodItem.CreatedByUserID, &foodItem.CreatedAt, &foodItem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		foodItems = append(foodItems, &foodItem)
	}
	return foodItems, nil
}

// Count returns the total number of food items matching the filters
func (r *foodItemRepository) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	query := r.db.Builder.Select("COUNT(*)").
		From("food_items")

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

// CountBySearch returns the total number of food items matching the search query
func (r *foodItemRepository) CountBySearch(ctx context.Context, query string) (int64, error) {
	sqlQuery, args, err := r.db.Builder.Select("COUNT(*)").
		From("food_items").
		Where(squirrel.Or{
			squirrel.Like{"name": "%" + query + "%"},
			squirrel.Like{"brand_name": "%" + query + "%"},
		}).
		ToSql()
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

// List returns a paginated list of food items matching the filters
func (r *foodItemRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.FoodItem, error) {
	query := r.db.Builder.Select("id", "name", "brand_name", "barcode_upc", "serving_size_default_qty", "serving_size_default_unit", "calories_per_default_serving", "protein_grams_per_default_serving", "fat_grams_per_default_serving", "carbs_grams_per_default_serving", "fiber_grams_per_default_serving", "sugar_grams_per_default_serving", "saturated_fat_grams_per_default_serving", "trans_fat_grams_per_default_serving", "cholesterol_mg_per_default_serving", "sodium_mg_per_default_serving", "potassium_mg_per_default_serving", "vitamin_a_mcg_per_default_serving", "vitamin_c_mg_per_default_serving", "calcium_mg_per_default_serving", "iron_mg_per_default_serving", "source", "is_verified", "created_by_user_id", "created_at", "updated_at").
		From("food_items")

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

	var foodItems []*entity.FoodItem
	for rows.Next() {
		var foodItem entity.FoodItem
		err := rows.Scan(
			&foodItem.ID, &foodItem.Name, &foodItem.BrandName, &foodItem.BarcodeUPC, &foodItem.ServingSizeDefaultQty, &foodItem.ServingSizeDefaultUnit, &foodItem.CaloriesPerDefaultServing, &foodItem.ProteinGramsPerDefaultServing, &foodItem.FatGramsPerDefaultServing, &foodItem.CarbsGramsPerDefaultServing, &foodItem.FiberGramsPerDefaultServing, &foodItem.SugarGramsPerDefaultServing, &foodItem.SaturatedFatGramsPerDefaultServing, &foodItem.TransFatGramsPerDefaultServing, &foodItem.CholesterolMgPerDefaultServing, &foodItem.SodiumMgPerDefaultServing, &foodItem.PotassiumMgPerDefaultServing, &foodItem.VitaminAMcgPerDefaultServing, &foodItem.VitaminCMgPerDefaultServing, &foodItem.CalciumMgPerDefaultServing, &foodItem.IronMgPerDefaultServing, &foodItem.Source, &foodItem.IsVerified, &foodItem.CreatedByUserID, &foodItem.CreatedAt, &foodItem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		foodItems = append(foodItems, &foodItem)
	}

	return foodItems, nil
}
