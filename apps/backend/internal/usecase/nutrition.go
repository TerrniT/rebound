package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/repository"
)

// NutritionUseCase handles nutrition-related business logic
type NutritionUseCase struct {
	nutritionRepo repository.NutritionRepository
}

// NewNutritionUseCase creates a new instance of NutritionUseCase
func NewNutritionUseCase(nutritionRepo repository.NutritionRepository) *NutritionUseCase {
	return &NutritionUseCase{
		nutritionRepo: nutritionRepo,
	}
}

// CreateNutritionGoals creates new nutrition goals
func (uc *NutritionUseCase) CreateNutritionGoals(ctx context.Context, goals *entity.UserNutritionGoal) (*entity.UserNutritionGoal, error) {
	// Generate new ID and timestamps
	goals.ID = uuid.New().String()
	now := time.Now()
	goals.CreatedAt = now
	goals.UpdatedAt = now

	return uc.nutritionRepo.CreateNutritionGoals(ctx, goals)
}

// GetNutritionGoals retrieves nutrition goals by ID
func (uc *NutritionUseCase) GetNutritionGoals(ctx context.Context, id string) (*entity.UserNutritionGoal, error) {
	return uc.nutritionRepo.GetNutritionGoalsByID(ctx, id)
}

// UpdateNutritionGoals updates existing nutrition goals
func (uc *NutritionUseCase) UpdateNutritionGoals(ctx context.Context, goals *entity.UserNutritionGoal) error {
	goals.UpdatedAt = time.Now()
	return uc.nutritionRepo.UpdateNutritionGoals(ctx, goals)
}

// DeleteNutritionGoals deletes nutrition goals
func (uc *NutritionUseCase) DeleteNutritionGoals(ctx context.Context, id string) error {
	return uc.nutritionRepo.DeleteNutritionGoals(ctx, id)
}

// GetActiveNutritionGoals retrieves the active nutrition goals for a user
func (uc *NutritionUseCase) GetActiveNutritionGoals(ctx context.Context, userID string) (*entity.UserNutritionGoal, error) {
	return uc.nutritionRepo.GetActiveNutritionGoals(ctx, userID)
}

// GetNutritionGoalsHistory retrieves nutrition goals history for a user
func (uc *NutritionUseCase) GetNutritionGoalsHistory(ctx context.Context, userID string, page, pageSize int) ([]*entity.UserNutritionGoal, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = 10 // Default page size
	}
	if pageSize > 100 {
		pageSize = 100 // Maximum page size
	}

	offset := (page - 1) * pageSize
	return uc.nutritionRepo.GetNutritionGoalsHistory(ctx, userID, pageSize, offset)
}

// CreateBiometrics creates new biometrics entry
func (uc *NutritionUseCase) CreateBiometrics(ctx context.Context, biometrics *entity.UserBiometric) (*entity.UserBiometric, error) {
	// Generate new ID and timestamp
	biometrics.ID = uuid.New().String()
	biometrics.CreatedAt = time.Now()

	return uc.nutritionRepo.CreateBiometrics(ctx, biometrics)
}

// GetBiometrics retrieves biometrics by ID
func (uc *NutritionUseCase) GetBiometrics(ctx context.Context, id string) (*entity.UserBiometric, error) {
	return uc.nutritionRepo.GetBiometricsByID(ctx, id)
}

// UpdateBiometrics updates existing biometrics
func (uc *NutritionUseCase) UpdateBiometrics(ctx context.Context, biometrics *entity.UserBiometric) error {
	return uc.nutritionRepo.UpdateBiometrics(ctx, biometrics)
}

// DeleteBiometrics deletes biometrics entry
func (uc *NutritionUseCase) DeleteBiometrics(ctx context.Context, id string) error {
	return uc.nutritionRepo.DeleteBiometrics(ctx, id)
}

// GetUserBiometricsHistory retrieves biometrics history for a user
func (uc *NutritionUseCase) GetUserBiometricsHistory(ctx context.Context, userID string, page, pageSize int) ([]*entity.UserBiometric, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = 10 // Default page size
	}
	if pageSize > 100 {
		pageSize = 100 // Maximum page size
	}

	offset := (page - 1) * pageSize
	return uc.nutritionRepo.GetUserBiometricsHistory(ctx, userID, pageSize, offset)
}

// GetLatestBiometrics retrieves the most recent biometrics for a user
func (uc *NutritionUseCase) GetLatestBiometrics(ctx context.Context, userID string) (*entity.UserBiometric, error) {
	return uc.nutritionRepo.GetLatestBiometrics(ctx, userID)
}
