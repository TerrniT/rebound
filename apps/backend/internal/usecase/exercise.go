package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/repository"
)

// ExerciseUseCase represents the exercise use case
type ExerciseUseCase struct {
	repo   repository.ExerciseRepository
	config Config
}

// NewExerciseUseCase creates a new instance of ExerciseUseCase
func NewExerciseUseCase(r repository.ExerciseRepository, config Config) *ExerciseUseCase {
	return &ExerciseUseCase{
		repo:   r,
		config: config,
	}
}

// CreateExercise creates a new exercise
func (uc *ExerciseUseCase) CreateExercise(ctx context.Context, exercise *entity.Exercise) (*entity.Exercise, error) {
	exercise.ID = uuid.New().String()
	exercise.CreatedAt = time.Now()
	exercise.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, exercise)
}

// GetExercise retrieves an exercise by its ID
func (uc *ExerciseUseCase) GetExercise(ctx context.Context, exerciseID string) (*entity.Exercise, error) {
	return uc.repo.GetByID(ctx, exerciseID)
}

// ListExercises returns a paginated list of exercises matching the filters
func (uc *ExerciseUseCase) ListExercises(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.Exercise, int64, error) {
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

	// Get exercises
	items, err := uc.repo.List(ctx, filters, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// UpdateExercise updates an existing exercise
func (uc *ExerciseUseCase) UpdateExercise(ctx context.Context, exercise *entity.Exercise) error {
	exercise.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, exercise)
}

// DeleteExercise deletes an exercise
func (uc *ExerciseUseCase) DeleteExercise(ctx context.Context, exerciseID string) error {
	return uc.repo.Delete(ctx, exerciseID)
}

// SearchExercises searches for exercises by name or description
func (uc *ExerciseUseCase) SearchExercises(ctx context.Context, query string, page, pageSize int) ([]*entity.Exercise, int64, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = uc.config.DefaultPageSize
	}
	if pageSize > uc.config.MaxPageSize {
		pageSize = uc.config.MaxPageSize
	}

	// Get total count
	total, err := uc.repo.Count(ctx, map[string]interface{}{
		"exercise_name": query,
	})
	if err != nil {
		return nil, 0, err
	}

	// Search exercises
	items, err := uc.repo.Search(ctx, query, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// GetUserExercises retrieves exercises created by a specific user
func (uc *ExerciseUseCase) GetUserExercises(ctx context.Context, userID string, page, pageSize int) ([]*entity.Exercise, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = uc.config.DefaultPageSize
	}
	if pageSize > uc.config.MaxPageSize {
		pageSize = uc.config.MaxPageSize
	}

	return uc.repo.GetByUserID(ctx, userID, pageSize, (page-1)*pageSize)
}
