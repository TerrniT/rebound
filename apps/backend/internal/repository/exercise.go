package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

// ExerciseRepository defines the interface for exercise-related database operations
type ExerciseRepository interface {
	Create(ctx context.Context, exercise *entity.Exercise) (*entity.Exercise, error)
	GetByID(ctx context.Context, id string) (*entity.Exercise, error)
	Update(ctx context.Context, exercise *entity.Exercise) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.Exercise, error)
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*entity.Exercise, error)
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.Exercise, error)
}

// exerciseRepository implements ExerciseRepository
type exerciseRepository struct {
	db *postgres.Postgres
}

// NewExerciseRepository creates a new instance of ExerciseRepository
func NewExerciseRepository(db *postgres.Postgres) ExerciseRepository {
	return &exerciseRepository{db: db}
}

// Create creates a new exercise in the database
func (r *exerciseRepository) Create(ctx context.Context, exercise *entity.Exercise) (*entity.Exercise, error) {
	query, args, err := r.db.Builder.Insert("exercises").
		Columns("exercise_id", "exercise_name", "description", "muscle_group_primary", "muscle_groups_secondary", "equipment_required", "difficulty_level", "video_url", "image_url_thumbnail", "image_url_main", "exercise_type", "created_by_user_id", "is_public", "created_at", "updated_at").
		Values(exercise.ID, exercise.Name, exercise.Description, exercise.MuscleGroupPrimary, exercise.MuscleGroupsSecondary, exercise.EquipmentRequired, exercise.DifficultyLevel, exercise.VideoURL, exercise.ImageURLThumbnail, exercise.ImageURLMain, exercise.Type, exercise.CreatedByUserID, exercise.IsPublic, exercise.CreatedAt, exercise.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return exercise, nil
}

// GetByID retrieves an exercise by its ID
func (r *exerciseRepository) GetByID(ctx context.Context, id string) (*entity.Exercise, error) {
	query, args, err := r.db.Builder.Select("exercise_id", "exercise_name", "description", "muscle_group_primary", "muscle_groups_secondary", "equipment_required", "difficulty_level", "video_url", "image_url_thumbnail", "image_url_main", "exercise_type", "created_by_user_id", "is_public", "created_at", "updated_at").
		From("exercises").
		Where(squirrel.Eq{"exercise_id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var exercise entity.Exercise
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&exercise.ID, &exercise.Name, &exercise.Description, &exercise.MuscleGroupPrimary, &exercise.MuscleGroupsSecondary, &exercise.EquipmentRequired, &exercise.DifficultyLevel, &exercise.VideoURL, &exercise.ImageURLThumbnail, &exercise.ImageURLMain, &exercise.Type, &exercise.CreatedByUserID, &exercise.IsPublic, &exercise.CreatedAt, &exercise.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &exercise, nil
}

// Update updates an existing exercise in the database
func (r *exerciseRepository) Update(ctx context.Context, exercise *entity.Exercise) error {
	query, args, err := r.db.Builder.Update("exercises").
		Set("exercise_name", exercise.Name).
		Set("description", exercise.Description).
		Set("muscle_group_primary", exercise.MuscleGroupPrimary).
		Set("muscle_groups_secondary", exercise.MuscleGroupsSecondary).
		Set("equipment_required", exercise.EquipmentRequired).
		Set("difficulty_level", exercise.DifficultyLevel).
		Set("video_url", exercise.VideoURL).
		Set("image_url_thumbnail", exercise.ImageURLThumbnail).
		Set("image_url_main", exercise.ImageURLMain).
		Set("exercise_type", exercise.Type).
		Set("created_by_user_id", exercise.CreatedByUserID).
		Set("is_public", exercise.IsPublic).
		Set("updated_at", exercise.UpdatedAt).
		Where(squirrel.Eq{"exercise_id": exercise.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// Delete deletes an exercise from the database
func (r *exerciseRepository) Delete(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("exercises").
		Where(squirrel.Eq{"exercise_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// List returns a paginated list of exercises matching the filters
func (r *exerciseRepository) List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.Exercise, error) {
	query := r.db.Builder.Select("exercise_id", "exercise_name", "description", "muscle_group_primary", "muscle_groups_secondary", "equipment_required", "difficulty_level", "video_url", "image_url_thumbnail", "image_url_main", "exercise_type", "created_by_user_id", "is_public", "created_at", "updated_at").
		From("exercises")

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

	var exercises []*entity.Exercise
	for rows.Next() {
		var exercise entity.Exercise
		err := rows.Scan(
			&exercise.ID, &exercise.Name, &exercise.Description, &exercise.MuscleGroupPrimary, &exercise.MuscleGroupsSecondary, &exercise.EquipmentRequired, &exercise.DifficultyLevel, &exercise.VideoURL, &exercise.ImageURLThumbnail, &exercise.ImageURLMain, &exercise.Type, &exercise.CreatedByUserID, &exercise.IsPublic, &exercise.CreatedAt, &exercise.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}

	return exercises, nil
}

// Count returns the total number of exercises matching the filters
func (r *exerciseRepository) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	query := r.db.Builder.Select("COUNT(*)").
		From("exercises")

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

// Search searches for exercises by name or description
func (r *exerciseRepository) Search(ctx context.Context, query string, limit, offset int) ([]*entity.Exercise, error) {
	sqlQuery, args, err := r.db.Builder.Select("exercise_id", "exercise_name", "description", "muscle_group_primary", "muscle_groups_secondary", "equipment_required", "difficulty_level", "video_url", "image_url_thumbnail", "image_url_main", "exercise_type", "created_by_user_id", "is_public", "created_at", "updated_at").
		From("exercises").
		Where(squirrel.Or{
			squirrel.Like{"exercise_name": "%" + query + "%"},
			squirrel.Like{"description": "%" + query + "%"},
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
	var exercises []*entity.Exercise
	for rows.Next() {
		var exercise entity.Exercise
		err := rows.Scan(
			&exercise.ID, &exercise.Name, &exercise.Description, &exercise.MuscleGroupPrimary, &exercise.MuscleGroupsSecondary, &exercise.EquipmentRequired, &exercise.DifficultyLevel, &exercise.VideoURL, &exercise.ImageURLThumbnail, &exercise.ImageURLMain, &exercise.Type, &exercise.CreatedByUserID, &exercise.IsPublic, &exercise.CreatedAt, &exercise.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}
	return exercises, nil
}

// GetByUserID retrieves exercises created by a specific user
func (r *exerciseRepository) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.Exercise, error) {
	query, args, err := r.db.Builder.Select("exercise_id", "exercise_name", "description", "muscle_group_primary", "muscle_groups_secondary", "equipment_required", "difficulty_level", "video_url", "image_url_thumbnail", "image_url_main", "exercise_type", "created_by_user_id", "is_public", "created_at", "updated_at").
		From("exercises").
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
	var exercises []*entity.Exercise
	for rows.Next() {
		var exercise entity.Exercise
		err := rows.Scan(
			&exercise.ID, &exercise.Name, &exercise.Description, &exercise.MuscleGroupPrimary, &exercise.MuscleGroupsSecondary, &exercise.EquipmentRequired, &exercise.DifficultyLevel, &exercise.VideoURL, &exercise.ImageURLThumbnail, &exercise.ImageURLMain, &exercise.Type, &exercise.CreatedByUserID, &exercise.IsPublic, &exercise.CreatedAt, &exercise.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}
	return exercises, nil
}
