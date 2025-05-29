package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/repository"
)

// MealUseCase handles meal-related business logic
type MealUseCase struct {
	mealRepo repository.MealRepository
}

// NewMealUseCase creates a new instance of MealUseCase
func NewMealUseCase(mealRepo repository.MealRepository) *MealUseCase {
	return &MealUseCase{
		mealRepo: mealRepo,
	}
}

// CreateMeal creates a new meal
func (uc *MealUseCase) CreateMeal(ctx context.Context, meal *entity.UserMeal) (*entity.UserMeal, error) {
	// Generate new ID and timestamps
	meal.ID = uuid.New().String()
	now := time.Now()
	meal.CreatedAt = now
	meal.UpdatedAt = now

	return uc.mealRepo.Create(ctx, meal)
}

// GetMeal retrieves a meal by its ID
func (uc *MealUseCase) GetMeal(ctx context.Context, id string) (*entity.UserMeal, error) {
	return uc.mealRepo.GetByID(ctx, id)
}

// ListMeals returns a paginated list of meals
func (uc *MealUseCase) ListMeals(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.UserMeal, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = 10 // Default page size
	}
	if pageSize > 100 {
		pageSize = 100 // Maximum page size
	}

	return uc.mealRepo.List(ctx, filters, page, pageSize)
}

// UpdateMeal updates an existing meal
func (uc *MealUseCase) UpdateMeal(ctx context.Context, meal *entity.UserMeal) error {
	meal.UpdatedAt = time.Now()
	return uc.mealRepo.Update(ctx, meal)
}

// DeleteMeal deletes a meal by its ID
func (uc *MealUseCase) DeleteMeal(ctx context.Context, id string) error {
	return uc.mealRepo.Delete(ctx, id)
}

// GetUserMeals retrieves meals for a specific user
func (uc *MealUseCase) GetUserMeals(ctx context.Context, userID string, page, pageSize int) ([]*entity.UserMeal, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = 10 // Default page size
	}
	if pageSize > 100 {
		pageSize = 100 // Maximum page size
	}

	offset := (page - 1) * pageSize
	return uc.mealRepo.GetByUserID(ctx, userID, pageSize, offset)
}

// AddFoodItemToMeal adds a new food item to a meal
func (uc *MealUseCase) AddFoodItemToMeal(ctx context.Context, foodItem *entity.MealFoodItem) error {
	// Generate new ID and timestamp
	foodItem.ID = uuid.New().String()
	foodItem.LoggedAt = time.Now()

	return uc.mealRepo.AddFoodItem(ctx, foodItem)
}

// GetMealFoodItems retrieves all food items for a meal
func (uc *MealUseCase) GetMealFoodItems(ctx context.Context, mealID string) ([]*entity.MealFoodItem, error) {
	return uc.mealRepo.GetFoodItems(ctx, mealID)
}

// UpdateMealFoodItem updates an existing food item
func (uc *MealUseCase) UpdateMealFoodItem(ctx context.Context, foodItem *entity.MealFoodItem) error {
	return uc.mealRepo.UpdateFoodItem(ctx, foodItem)
}

// DeleteMealFoodItem deletes a food item
func (uc *MealUseCase) DeleteMealFoodItem(ctx context.Context, foodItemID string) error {
	return uc.mealRepo.DeleteFoodItem(ctx, foodItemID)
}
