package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/repository"
)

// WorkoutSessionUseCase represents the workout session use case
type WorkoutSessionUseCase struct {
	repo   repository.WorkoutSessionRepository
	config Config
}

// NewWorkoutSessionUseCase creates a new instance of WorkoutSessionUseCase
func NewWorkoutSessionUseCase(r repository.WorkoutSessionRepository, config Config) *WorkoutSessionUseCase {
	return &WorkoutSessionUseCase{
		repo:   r,
		config: config,
	}
}

// CreateWorkoutSession creates a new workout session
func (uc *WorkoutSessionUseCase) CreateWorkoutSession(ctx context.Context, session *entity.UserWorkoutSession) (*entity.UserWorkoutSession, error) {
	session.ID = uuid.New().String()
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, session)
}

// GetWorkoutSession retrieves a workout session by its ID
func (uc *WorkoutSessionUseCase) GetWorkoutSession(ctx context.Context, sessionID string) (*entity.UserWorkoutSession, error) {
	return uc.repo.GetByID(ctx, sessionID)
}

// ListWorkoutSessions returns a paginated list of workout sessions matching the filters
func (uc *WorkoutSessionUseCase) ListWorkoutSessions(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.UserWorkoutSession, int64, error) {
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

	// Get workout sessions
	items, err := uc.repo.List(ctx, filters, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// UpdateWorkoutSession updates an existing workout session
func (uc *WorkoutSessionUseCase) UpdateWorkoutSession(ctx context.Context, session *entity.UserWorkoutSession) error {
	session.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, session)
}

// DeleteWorkoutSession deletes a workout session
func (uc *WorkoutSessionUseCase) DeleteWorkoutSession(ctx context.Context, sessionID string) error {
	return uc.repo.Delete(ctx, sessionID)
}

// GetUserWorkoutSessions retrieves workout sessions for a specific user
func (uc *WorkoutSessionUseCase) GetUserWorkoutSessions(ctx context.Context, userID string, page, pageSize int) ([]*entity.UserWorkoutSession, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = uc.config.DefaultPageSize
	}
	if pageSize > uc.config.MaxPageSize {
		pageSize = uc.config.MaxPageSize
	}

	return uc.repo.GetByUserID(ctx, userID, pageSize, (page-1)*pageSize)
}

// AddSessionLog adds a new log entry to a workout session
func (uc *WorkoutSessionUseCase) AddSessionLog(ctx context.Context, log *entity.UserWorkoutSessionLog) error {
	log.ID = uuid.New().String()
	log.LoggedAt = time.Now()
	return uc.repo.AddLog(ctx, log)
}

// GetSessionLogs retrieves all logs for a workout session
func (uc *WorkoutSessionUseCase) GetSessionLogs(ctx context.Context, sessionID string) ([]*entity.UserWorkoutSessionLog, error) {
	return uc.repo.GetLogs(ctx, sessionID)
}

// UpdateSessionLog updates an existing log entry
func (uc *WorkoutSessionUseCase) UpdateSessionLog(ctx context.Context, log *entity.UserWorkoutSessionLog) error {
	return uc.repo.UpdateLog(ctx, log)
}

// DeleteSessionLog deletes a log entry
func (uc *WorkoutSessionUseCase) DeleteSessionLog(ctx context.Context, logID string) error {
	return uc.repo.DeleteLog(ctx, logID)
}
