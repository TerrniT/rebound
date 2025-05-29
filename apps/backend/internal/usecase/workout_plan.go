package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/repository"
)

// WorkoutPlanUseCase represents the workout plan use case
type WorkoutPlanUseCase struct {
	repo   repository.WorkoutPlanRepository
	config Config
}

// NewWorkoutPlanUseCase creates a new instance of WorkoutPlanUseCase
func NewWorkoutPlanUseCase(r repository.WorkoutPlanRepository, config Config) *WorkoutPlanUseCase {
	return &WorkoutPlanUseCase{
		repo:   r,
		config: config,
	}
}

// CreateWorkoutPlan creates a new workout plan
func (uc *WorkoutPlanUseCase) CreateWorkoutPlan(ctx context.Context, plan *entity.WorkoutPlan) (*entity.WorkoutPlan, error) {
	plan.ID = uuid.New().String()
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, plan)
}

// GetWorkoutPlan retrieves a workout plan by its ID
func (uc *WorkoutPlanUseCase) GetWorkoutPlan(ctx context.Context, planID string) (*entity.WorkoutPlan, error) {
	return uc.repo.GetByID(ctx, planID)
}

// ListWorkoutPlans returns a paginated list of workout plans matching the filters
func (uc *WorkoutPlanUseCase) ListWorkoutPlans(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.WorkoutPlan, int64, error) {
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

	// Get workout plans
	items, err := uc.repo.List(ctx, filters, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// UpdateWorkoutPlan updates an existing workout plan
func (uc *WorkoutPlanUseCase) UpdateWorkoutPlan(ctx context.Context, plan *entity.WorkoutPlan) error {
	plan.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, plan)
}

// DeleteWorkoutPlan deletes a workout plan
func (uc *WorkoutPlanUseCase) DeleteWorkoutPlan(ctx context.Context, planID string) error {
	return uc.repo.Delete(ctx, planID)
}

// GetUserWorkoutPlans retrieves workout plans created by a specific user
func (uc *WorkoutPlanUseCase) GetUserWorkoutPlans(ctx context.Context, userID string, page, pageSize int) ([]*entity.WorkoutPlan, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = uc.config.DefaultPageSize
	}
	if pageSize > uc.config.MaxPageSize {
		pageSize = uc.config.MaxPageSize
	}

	return uc.repo.GetByUserID(ctx, userID, pageSize, (page-1)*pageSize)
}

// AddExerciseToPlan adds an exercise to a workout plan
func (uc *WorkoutPlanUseCase) AddExerciseToPlan(ctx context.Context, planExercise *entity.WorkoutPlanExercise) error {
	planExercise.ID = uuid.New().String()
	return uc.repo.AddExercise(ctx, planExercise)
}

// RemoveExerciseFromPlan removes an exercise from a workout plan
func (uc *WorkoutPlanUseCase) RemoveExerciseFromPlan(ctx context.Context, planID, exerciseID string) error {
	return uc.repo.RemoveExercise(ctx, planID, exerciseID)
}

// GetPlanExercises retrieves all exercises in a workout plan
func (uc *WorkoutPlanUseCase) GetPlanExercises(ctx context.Context, planID string) ([]*entity.WorkoutPlanExercise, error) {
	return uc.repo.GetExercises(ctx, planID)
}
