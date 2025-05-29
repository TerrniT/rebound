package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

// WorkoutPlanRepository defines the interface for workout plan-related database operations
type WorkoutPlanRepository interface {
	Create(ctx context.Context, plan *entity.WorkoutPlan) (*entity.WorkoutPlan, error)
	GetByID(ctx context.Context, id string) (*entity.WorkoutPlan, error)
	Update(ctx context.Context, plan *entity.WorkoutPlan) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.WorkoutPlan, error)
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.WorkoutPlan, error)
	AddExercise(ctx context.Context, planExercise *entity.WorkoutPlanExercise) error
	RemoveExercise(ctx context.Context, planID, exerciseID string) error
	GetExercises(ctx context.Context, planID string) ([]*entity.WorkoutPlanExercise, error)
}

// workoutPlanRepository implements WorkoutPlanRepository
type workoutPlanRepository struct {
	db *postgres.Postgres
}

// NewWorkoutPlanRepository creates a new instance of WorkoutPlanRepository
func NewWorkoutPlanRepository(db *postgres.Postgres) WorkoutPlanRepository {
	return &workoutPlanRepository{db: db}
}

// Create creates a new workout plan in the database
func (r *workoutPlanRepository) Create(ctx context.Context, plan *entity.WorkoutPlan) (*entity.WorkoutPlan, error) {
	query, args, err := r.db.Builder.Insert("workout_plans").
		Columns("plan_id", "user_id", "plan_name", "description", "plan_type", "difficulty_level", "duration_estimate_minutes", "frequency_per_week", "is_public", "cover_image_url", "created_at", "updated_at").
		Values(plan.ID, plan.UserID, plan.Name, plan.Description, plan.DifficultyLevel, plan.DurationEstimate, plan.FrequencyPerWeek, plan.IsPublic, plan.CoverImageURL, plan.CreatedAt, plan.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

// GetByID retrieves a workout plan by its ID
func (r *workoutPlanRepository) GetByID(ctx context.Context, id string) (*entity.WorkoutPlan, error) {
	query, args, err := r.db.Builder.Select("plan_id", "user_id", "plan_name", "description", "plan_type", "difficulty_level", "duration_estimate_minutes", "frequency_per_week", "is_public", "cover_image_url", "created_at", "updated_at").
		From("workout_plans").
		Where(squirrel.Eq{"plan_id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var plan entity.WorkoutPlan
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&plan.ID, &plan.UserID, &plan.Name, &plan.Description, &plan.DifficultyLevel, &plan.DurationEstimate, &plan.FrequencyPerWeek, &plan.IsPublic, &plan.CoverImageURL, &plan.CreatedAt, &plan.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

// Update updates an existing workout plan in the database
func (r *workoutPlanRepository) Update(ctx context.Context, plan *entity.WorkoutPlan) error {
	query, args, err := r.db.Builder.Update("workout_plans").
		Set("user_id", plan.UserID).
		Set("plan_name", plan.Name).
		Set("description", plan.Description).
		// Set("plan_type", plan.PlanType).
		Set("difficulty_level", plan.DifficultyLevel).
		Set("duration_estimate", plan.DurationEstimate).
		Set("frequency_per_week", plan.FrequencyPerWeek).
		Set("is_public", plan.IsPublic).
		Set("cover_image_url", plan.CoverImageURL).
		Set("updated_at", plan.UpdatedAt).
		Where(squirrel.Eq{"plan_id": plan.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// Delete deletes a workout plan from the database
func (r *workoutPlanRepository) Delete(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("workout_plans").
		Where(squirrel.Eq{"plan_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// List returns a paginated list of workout plans matching the filters
func (r *workoutPlanRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.WorkoutPlan, error) {
	query := r.db.Builder.Select("plan_id", "user_id", "plan_name", "description", "plan_type", "difficulty_level", "duration_estimate_minutes", "frequency_per_week", "is_public", "cover_image_url", "created_at", "updated_at").
		From("workout_plans")

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

	var plans []*entity.WorkoutPlan
	for rows.Next() {
		var plan entity.WorkoutPlan
		err := rows.Scan(
			&plan.ID, &plan.UserID, &plan.Name, &plan.Description, &plan.DifficultyLevel, &plan.DurationEstimate, &plan.FrequencyPerWeek, &plan.IsPublic, &plan.CoverImageURL, &plan.CreatedAt, &plan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, &plan)
	}

	return plans, nil
}

// Count returns the total number of workout plans matching the filters
func (r *workoutPlanRepository) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	query := r.db.Builder.Select("COUNT(*)").
		From("workout_plans")

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

// GetByUserID retrieves workout plans created by a specific user
func (r *workoutPlanRepository) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.WorkoutPlan, error) {
	query, args, err := r.db.Builder.Select("plan_id", "user_id", "plan_name", "description", "plan_type", "difficulty_level", "duration_estimate", "frequency_per_week", "is_public", "cover_image_url", "created_at", "updated_at").
		From("workout_plans").
		Where(squirrel.Eq{"user_id": userID}).
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
	var plans []*entity.WorkoutPlan
	for rows.Next() {
		var plan entity.WorkoutPlan
		err := rows.Scan(
			&plan.ID, &plan.UserID, &plan.Name, &plan.Description, &plan.DifficultyLevel, &plan.DurationEstimate, &plan.FrequencyPerWeek, &plan.IsPublic, &plan.CoverImageURL, &plan.CreatedAt, &plan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, &plan)
	}
	return plans, nil
}

// AddExercise adds an exercise to a workout plan
func (r *workoutPlanRepository) AddExercise(ctx context.Context, planExercise *entity.WorkoutPlanExercise) error {
	query, args, err := r.db.Builder.Insert("workout_plan_exercises").
		Columns("plan_exercise_id", "plan_id", "exercise_id", "day_of_week", "day_number", "exercise_order", "sets", "reps_min", "reps_max", "reps_target", "duration_seconds", "rest_period_seconds", "notes").
		Values(planExercise.ID, planExercise.PlanID, planExercise.ExerciseID, planExercise.DayOfWeek, planExercise.DayNumber, planExercise.ExerciseOrder, planExercise.Sets, planExercise.RepsMin, planExercise.RepsMax, planExercise.RepsTarget, planExercise.DurationSeconds, planExercise.RestPeriodSeconds, planExercise.Notes).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// RemoveExercise removes an exercise from a workout plan
func (r *workoutPlanRepository) RemoveExercise(ctx context.Context, planID, exerciseID string) error {
	query, args, err := r.db.Builder.Delete("workout_plan_exercises").
		Where(squirrel.Eq{"plan_id": planID, "exercise_id": exerciseID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// GetExercises retrieves all exercises in a workout plan
func (r *workoutPlanRepository) GetExercises(ctx context.Context, planID string) ([]*entity.WorkoutPlanExercise, error) {
	query, args, err := r.db.Builder.Select("plan_exercise_id", "plan_id", "exercise_id", "day_of_week", "day_number", "exercise_order", "sets", "reps_min", "reps_max", "reps_target", "duration_seconds", "rest_period_seconds", "notes").
		From("workout_plan_exercises").
		Where(squirrel.Eq{"plan_id": planID}).
		OrderBy("exercise_order").
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var exercises []*entity.WorkoutPlanExercise
	for rows.Next() {
		var exercise entity.WorkoutPlanExercise
		err := rows.Scan(
			&exercise.ID, &exercise.PlanID, &exercise.ExerciseID, &exercise.DayOfWeek, &exercise.DayNumber, &exercise.ExerciseOrder, &exercise.Sets, &exercise.RepsMin, &exercise.RepsMax, &exercise.RepsTarget, &exercise.DurationSeconds, &exercise.RestPeriodSeconds, &exercise.Notes,
		)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}
	return exercises, nil
}
