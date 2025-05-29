package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

// WorkoutSessionRepository defines the interface for workout session-related database operations
type WorkoutSessionRepository interface {
	Create(ctx context.Context, session *entity.UserWorkoutSession) (*entity.UserWorkoutSession, error)
	GetByID(ctx context.Context, id string) (*entity.UserWorkoutSession, error)
	Update(ctx context.Context, session *entity.UserWorkoutSession) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.UserWorkoutSession, error)
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.UserWorkoutSession, error)
	AddLog(ctx context.Context, log *entity.UserWorkoutSessionLog) error
	GetLogs(ctx context.Context, sessionID string) ([]*entity.UserWorkoutSessionLog, error)
	UpdateLog(ctx context.Context, log *entity.UserWorkoutSessionLog) error
	DeleteLog(ctx context.Context, logID string) error
}

// workoutSessionRepository implements WorkoutSessionRepository
type workoutSessionRepository struct {
	db *postgres.Postgres
}

// NewWorkoutSessionRepository creates a new instance of WorkoutSessionRepository
func NewWorkoutSessionRepository(db *postgres.Postgres) WorkoutSessionRepository {
	return &workoutSessionRepository{db: db}
}

// Create creates a new workout session in the database
func (r *workoutSessionRepository) Create(ctx context.Context, session *entity.UserWorkoutSession) (*entity.UserWorkoutSession, error) {
	query, args, err := r.db.Builder.Insert("user_workout_sessions").
		Columns("session_id", "user_id", "plan_id", "session_name", "scheduled_at", "started_at", "completed_at", "duration_minutes", "status", "notes", "location", "mood_rating", "perceived_exertion_rating", "created_at", "updated_at").
		Values(session.ID, session.UserID, session.PlanID, session.SessionName, session.ScheduledAt, session.StartedAt, session.CompletedAt, session.DurationMinutes, session.Status, session.Notes, session.Location, session.MoodRating, session.PerceivedExertionRating, session.CreatedAt, session.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetByID retrieves a workout session by its ID
func (r *workoutSessionRepository) GetByID(ctx context.Context, id string) (*entity.UserWorkoutSession, error) {
	query, args, err := r.db.Builder.Select("session_id", "user_id", "plan_id", "session_name", "scheduled_at", "started_at", "completed_at", "duration_minutes", "status", "notes", "location", "mood_rating", "perceived_exertion_rating", "created_at", "updated_at").
		From("user_workout_sessions").
		Where(squirrel.Eq{"session_id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var session entity.UserWorkoutSession
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&session.ID, &session.UserID, &session.PlanID, &session.SessionName, &session.ScheduledAt, &session.StartedAt, &session.CompletedAt, &session.DurationMinutes, &session.Status, &session.Notes, &session.Location, &session.MoodRating, &session.PerceivedExertionRating, &session.CreatedAt, &session.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// Update updates an existing workout session in the database
func (r *workoutSessionRepository) Update(ctx context.Context, session *entity.UserWorkoutSession) error {
	query, args, err := r.db.Builder.Update("user_workout_sessions").
		Set("user_id", session.UserID).
		Set("plan_id", session.PlanID).
		Set("session_name", session.SessionName).
		Set("scheduled_at", session.ScheduledAt).
		Set("started_at", session.StartedAt).
		Set("completed_at", session.CompletedAt).
		Set("duration_minutes", session.DurationMinutes).
		Set("status", session.Status).
		Set("notes", session.Notes).
		Set("location", session.Location).
		Set("mood_rating", session.MoodRating).
		Set("perceived_exertion_rating", session.PerceivedExertionRating).
		Set("updated_at", session.UpdatedAt).
		Where(squirrel.Eq{"session_id": session.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// Delete deletes a workout session from the database
func (r *workoutSessionRepository) Delete(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("user_workout_sessions").
		Where(squirrel.Eq{"session_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// List returns a paginated list of workout sessions matching the filters
func (r *workoutSessionRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.UserWorkoutSession, error) {
	query := r.db.Builder.Select("session_id", "user_id", "plan_id", "session_name", "scheduled_at", "started_at", "completed_at", "duration_minutes", "status", "notes", "location", "mood_rating", "perceived_exertion_rating", "created_at", "updated_at").
		From("user_workout_sessions")

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

	var sessions []*entity.UserWorkoutSession
	for rows.Next() {
		var session entity.UserWorkoutSession
		err := rows.Scan(
			&session.ID, &session.UserID, &session.PlanID, &session.SessionName, &session.ScheduledAt, &session.StartedAt, &session.CompletedAt, &session.DurationMinutes, &session.Status, &session.Notes, &session.Location, &session.MoodRating, &session.PerceivedExertionRating, &session.CreatedAt, &session.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, &session)
	}

	return sessions, nil
}

// Count returns the total number of workout sessions matching the filters
func (r *workoutSessionRepository) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	query := r.db.Builder.Select("COUNT(*)").
		From("user_workout_sessions")

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

// GetByUserID retrieves workout sessions for a specific user
func (r *workoutSessionRepository) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.UserWorkoutSession, error) {
	query, args, err := r.db.Builder.Select("session_id", "user_id", "plan_id", "session_name", "scheduled_at", "started_at", "completed_at", "duration_minutes", "status", "notes", "location", "mood_rating", "perceived_exertion_rating", "created_at", "updated_at").
		From("user_workout_sessions").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("scheduled_at DESC").
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
	var sessions []*entity.UserWorkoutSession
	for rows.Next() {
		var session entity.UserWorkoutSession
		err := rows.Scan(
			&session.ID, &session.UserID, &session.PlanID, &session.SessionName, &session.ScheduledAt, &session.StartedAt, &session.CompletedAt, &session.DurationMinutes, &session.Status, &session.Notes, &session.Location, &session.MoodRating, &session.PerceivedExertionRating, &session.CreatedAt, &session.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, &session)
	}
	return sessions, nil
}

// AddLog adds a new log entry to a workout session
func (r *workoutSessionRepository) AddLog(ctx context.Context, log *entity.UserWorkoutSessionLog) error {
	query, args, err := r.db.Builder.Insert("user_workout_session_logs").
		Columns("log_id", "session_id", "exercise_id", "plan_exercise_id", "set_number", "reps_completed", "weight_kg", "distance_km", "duration_seconds_completed", "rest_taken_seconds", "notes", "logged_at").
		Values(log.ID, log.SessionID, log.ExerciseID, log.PlanExerciseID, log.SetNumber, log.RepsCompleted, log.WeightKg, log.DistanceKm, log.DurationSecondsCompleted, log.RestTakenSeconds, log.Notes, log.LoggedAt).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// GetLogs retrieves all logs for a workout session
func (r *workoutSessionRepository) GetLogs(ctx context.Context, sessionID string) ([]*entity.UserWorkoutSessionLog, error) {
	query, args, err := r.db.Builder.Select("log_id", "session_id", "exercise_id", "plan_exercise_id", "set_number", "reps_completed", "weight_kg", "distance_km", "duration_seconds_completed", "rest_taken_seconds", "notes", "logged_at").
		From("user_workout_session_logs").
		Where(squirrel.Eq{"session_id": sessionID}).
		OrderBy("exercise_id", "set_number").
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var logs []*entity.UserWorkoutSessionLog
	for rows.Next() {
		var log entity.UserWorkoutSessionLog
		err := rows.Scan(
			&log.ID, &log.SessionID, &log.ExerciseID, &log.PlanExerciseID, &log.SetNumber, &log.RepsCompleted, &log.WeightKg, &log.DistanceKm, &log.DurationSecondsCompleted, &log.RestTakenSeconds, &log.Notes, &log.LoggedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}
	return logs, nil
}

// UpdateLog updates an existing log entry
func (r *workoutSessionRepository) UpdateLog(ctx context.Context, log *entity.UserWorkoutSessionLog) error {
	query, args, err := r.db.Builder.Update("user_workout_session_logs").
		Set("reps_completed", log.RepsCompleted).
		Set("weight_kg", log.WeightKg).
		Set("distance_km", log.DistanceKm).
		Set("duration_seconds_completed", log.DurationSecondsCompleted).
		Set("rest_taken_seconds", log.RestTakenSeconds).
		Set("notes", log.Notes).
		Where(squirrel.Eq{"log_id": log.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// DeleteLog deletes a log entry
func (r *workoutSessionRepository) DeleteLog(ctx context.Context, logID string) error {
	query, args, err := r.db.Builder.Delete("user_workout_session_logs").
		Where(squirrel.Eq{"log_id": logID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}
