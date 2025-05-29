package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

// NutritionRepository defines the interface for nutrition-related database operations
type NutritionRepository interface {
	// Nutrition Goals
	CreateNutritionGoals(ctx context.Context, goals *entity.UserNutritionGoal) (*entity.UserNutritionGoal, error)
	GetNutritionGoalsByID(ctx context.Context, id string) (*entity.UserNutritionGoal, error)
	UpdateNutritionGoals(ctx context.Context, goals *entity.UserNutritionGoal) error
	DeleteNutritionGoals(ctx context.Context, id string) error
	GetActiveNutritionGoals(ctx context.Context, userID string) (*entity.UserNutritionGoal, error)
	GetNutritionGoalsHistory(ctx context.Context, userID string, limit, offset int) ([]*entity.UserNutritionGoal, error)

	// Biometrics
	CreateBiometrics(ctx context.Context, biometrics *entity.UserBiometric) (*entity.UserBiometric, error)
	GetBiometricsByID(ctx context.Context, id string) (*entity.UserBiometric, error)
	UpdateBiometrics(ctx context.Context, biometrics *entity.UserBiometric) error
	DeleteBiometrics(ctx context.Context, id string) error
	GetUserBiometricsHistory(ctx context.Context, userID string, limit, offset int) ([]*entity.UserBiometric, error)
	GetLatestBiometrics(ctx context.Context, userID string) (*entity.UserBiometric, error)
}

// nutritionRepository implements NutritionRepository
type nutritionRepository struct {
	db *postgres.Postgres
}

// NewNutritionRepository creates a new instance of NutritionRepository
func NewNutritionRepository(db *postgres.Postgres) NutritionRepository {
	return &nutritionRepository{db: db}
}

// CreateNutritionGoals creates new nutrition goals
func (r *nutritionRepository) CreateNutritionGoals(ctx context.Context, goals *entity.UserNutritionGoal) (*entity.UserNutritionGoal, error) {
	query, args, err := r.db.Builder.Insert("user_nutrition_goals").
		Columns("nutrition_goals_id", "user_id", "goal_effective_date", "target_calories", "target_protein_grams", "target_fat_grams", "target_carbs_grams", "target_fiber_grams", "target_sugar_grams_limit", "notes", "is_active", "created_at", "updated_at").
		Values(goals.ID, goals.UserID, goals.GoalEffectiveDate, goals.TargetCalories, goals.TargetProteinGrams, goals.TargetFatGrams, goals.TargetCarbsGrams, goals.TargetFiberGrams, goals.TargetSugarGramsLimit, goals.Notes, goals.IsActive, goals.CreatedAt, goals.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return goals, nil
}

// GetNutritionGoalsByID retrieves nutrition goals by ID
func (r *nutritionRepository) GetNutritionGoalsByID(ctx context.Context, id string) (*entity.UserNutritionGoal, error) {
	query, args, err := r.db.Builder.Select("nutrition_goals_id", "user_id", "goal_effective_date", "target_calories", "target_protein_grams", "target_fat_grams", "target_carbs_grams", "target_fiber_grams", "target_sugar_grams_limit", "notes", "is_active", "created_at", "updated_at").
		From("user_nutrition_goals").
		Where(squirrel.Eq{"nutrition_goals_id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var goals entity.UserNutritionGoal
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&goals.ID, &goals.UserID, &goals.GoalEffectiveDate, &goals.TargetCalories, &goals.TargetProteinGrams, &goals.TargetFatGrams, &goals.TargetCarbsGrams, &goals.TargetFiberGrams, &goals.TargetSugarGramsLimit, &goals.Notes, &goals.IsActive, &goals.CreatedAt, &goals.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &goals, nil
}

// UpdateNutritionGoals updates existing nutrition goals
func (r *nutritionRepository) UpdateNutritionGoals(ctx context.Context, goals *entity.UserNutritionGoal) error {
	query, args, err := r.db.Builder.Update("user_nutrition_goals").
		Set("user_id", goals.UserID).
		Set("goal_effective_date", goals.GoalEffectiveDate).
		Set("target_calories", goals.TargetCalories).
		Set("target_protein_grams", goals.TargetProteinGrams).
		Set("target_fat_grams", goals.TargetFatGrams).
		Set("target_carbs_grams", goals.TargetCarbsGrams).
		Set("target_fiber_grams", goals.TargetFiberGrams).
		Set("target_sugar_grams_limit", goals.TargetSugarGramsLimit).
		Set("notes", goals.Notes).
		Set("is_active", goals.IsActive).
		Set("updated_at", goals.UpdatedAt).
		Where(squirrel.Eq{"nutrition_goals_id": goals.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// DeleteNutritionGoals deletes nutrition goals
func (r *nutritionRepository) DeleteNutritionGoals(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("user_nutrition_goals").
		Where(squirrel.Eq{"nutrition_goals_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// GetActiveNutritionGoals retrieves the active nutrition goals for a user
func (r *nutritionRepository) GetActiveNutritionGoals(ctx context.Context, userID string) (*entity.UserNutritionGoal, error) {
	query, args, err := r.db.Builder.Select("nutrition_goals_id", "user_id", "goal_effective_date", "target_calories", "target_protein_grams", "target_fat_grams", "target_carbs_grams", "target_fiber_grams", "target_sugar_grams_limit", "notes", "is_active", "created_at", "updated_at").
		From("user_nutrition_goals").
		Where(squirrel.Eq{"user_id": userID}).
		Where(squirrel.Eq{"is_active": true}).
		OrderBy("goal_effective_date DESC").
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}
	var goals entity.UserNutritionGoal
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&goals.ID, &goals.UserID, &goals.GoalEffectiveDate, &goals.TargetCalories, &goals.TargetProteinGrams, &goals.TargetFatGrams, &goals.TargetCarbsGrams, &goals.TargetFiberGrams, &goals.TargetSugarGramsLimit, &goals.Notes, &goals.IsActive, &goals.CreatedAt, &goals.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &goals, nil
}

// GetNutritionGoalsHistory retrieves nutrition goals history for a user
func (r *nutritionRepository) GetNutritionGoalsHistory(ctx context.Context, userID string, limit, offset int) ([]*entity.UserNutritionGoal, error) {
	query, args, err := r.db.Builder.Select("nutrition_goals_id", "user_id", "goal_effective_date", "target_calories", "target_protein_grams", "target_fat_grams", "target_carbs_grams", "target_fiber_grams", "target_sugar_grams_limit", "notes", "is_active", "created_at", "updated_at").
		From("user_nutrition_goals").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("goal_effective_date DESC").
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
	var goals []*entity.UserNutritionGoal
	for rows.Next() {
		var goal entity.UserNutritionGoal
		err := rows.Scan(
			&goal.ID, &goal.UserID, &goal.GoalEffectiveDate, &goal.TargetCalories, &goal.TargetProteinGrams, &goal.TargetFatGrams, &goal.TargetCarbsGrams, &goal.TargetFiberGrams, &goal.TargetSugarGramsLimit, &goal.Notes, &goal.IsActive, &goal.CreatedAt, &goal.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		goals = append(goals, &goal)
	}
	return goals, nil
}

// CreateBiometrics creates new biometrics entry
func (r *nutritionRepository) CreateBiometrics(ctx context.Context, biometrics *entity.UserBiometric) (*entity.UserBiometric, error) {
	query, args, err := r.db.Builder.Insert("user_biometrics").
		Columns("biometrics_id", "user_id", "log_date", "weight_kg", "height_cm", "body_fat_percentage", "waist_circumference_cm", "hip_circumference_cm", "chest_circumference_cm", "resting_heart_rate_bpm", "activity_level", "created_at").
		Values(biometrics.ID, biometrics.UserID, biometrics.LogDate, biometrics.WeightKg, biometrics.HeightCm, biometrics.BodyFatPercentage, biometrics.WaistCircumferenceCm, biometrics.HipCircumferenceCm, biometrics.ChestCircumferenceCm, biometrics.RestingHeartRateBpm, biometrics.ActivityLevel, biometrics.CreatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return biometrics, nil
}

// GetBiometricsByID retrieves biometrics by ID
func (r *nutritionRepository) GetBiometricsByID(ctx context.Context, id string) (*entity.UserBiometric, error) {
	query, args, err := r.db.Builder.Select("biometrics_id", "user_id", "log_date", "weight_kg", "height_cm", "body_fat_percentage", "waist_circumference_cm", "hip_circumference_cm", "chest_circumference_cm", "resting_heart_rate_bpm", "activity_level", "created_at").
		From("user_biometrics").
		Where(squirrel.Eq{"biometrics_id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var biometrics entity.UserBiometric
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&biometrics.ID, &biometrics.UserID, &biometrics.LogDate, &biometrics.WeightKg, &biometrics.HeightCm, &biometrics.BodyFatPercentage, &biometrics.WaistCircumferenceCm, &biometrics.HipCircumferenceCm, &biometrics.ChestCircumferenceCm, &biometrics.RestingHeartRateBpm, &biometrics.ActivityLevel, &biometrics.CreatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &biometrics, nil
}

// UpdateBiometrics updates existing biometrics
func (r *nutritionRepository) UpdateBiometrics(ctx context.Context, biometrics *entity.UserBiometric) error {
	query, args, err := r.db.Builder.Update("user_biometrics").
		Set("user_id", biometrics.UserID).
		Set("log_date", biometrics.LogDate).
		Set("weight_kg", biometrics.WeightKg).
		Set("height_cm", biometrics.HeightCm).
		Set("body_fat_percentage", biometrics.BodyFatPercentage).
		Set("waist_circumference_cm", biometrics.WaistCircumferenceCm).
		Set("hip_circumference_cm", biometrics.HipCircumferenceCm).
		Set("chest_circumference_cm", biometrics.ChestCircumferenceCm).
		Set("resting_heart_rate_bpm", biometrics.RestingHeartRateBpm).
		Set("activity_level", biometrics.ActivityLevel).
		Where(squirrel.Eq{"biometrics_id": biometrics.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// DeleteBiometrics deletes biometrics entry
func (r *nutritionRepository) DeleteBiometrics(ctx context.Context, id string) error {
	query, args, err := r.db.Builder.Delete("user_biometrics").
		Where(squirrel.Eq{"biometrics_id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Pool.Exec(ctx, query, args...)
	return err
}

// GetUserBiometricsHistory retrieves biometrics history for a user
func (r *nutritionRepository) GetUserBiometricsHistory(ctx context.Context, userID string, limit, offset int) ([]*entity.UserBiometric, error) {
	query, args, err := r.db.Builder.Select("biometrics_id", "user_id", "log_date", "weight_kg", "height_cm", "body_fat_percentage", "waist_circumference_cm", "hip_circumference_cm", "chest_circumference_cm", "resting_heart_rate_bpm", "activity_level", "created_at").
		From("user_biometrics").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("log_date DESC").
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
	var biometrics []*entity.UserBiometric
	for rows.Next() {
		var bio entity.UserBiometric
		err := rows.Scan(
			&bio.ID, &bio.UserID, &bio.LogDate, &bio.WeightKg, &bio.HeightCm, &bio.BodyFatPercentage, &bio.WaistCircumferenceCm, &bio.HipCircumferenceCm, &bio.ChestCircumferenceCm, &bio.RestingHeartRateBpm, &bio.ActivityLevel, &bio.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		biometrics = append(biometrics, &bio)
	}
	return biometrics, nil
}

// GetLatestBiometrics retrieves the most recent biometrics for a user
func (r *nutritionRepository) GetLatestBiometrics(ctx context.Context, userID string) (*entity.UserBiometric, error) {
	query, args, err := r.db.Builder.Select("biometrics_id", "user_id", "log_date", "weight_kg", "height_cm", "body_fat_percentage", "waist_circumference_cm", "hip_circumference_cm", "chest_circumference_cm", "resting_heart_rate_bpm", "activity_level", "created_at").
		From("user_biometrics").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("log_date DESC").
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}
	var biometrics entity.UserBiometric
	err = r.db.Pool.QueryRow(ctx, query, args...).Scan(
		&biometrics.ID, &biometrics.UserID, &biometrics.LogDate, &biometrics.WeightKg, &biometrics.HeightCm, &biometrics.BodyFatPercentage, &biometrics.WaistCircumferenceCm, &biometrics.HipCircumferenceCm, &biometrics.ChestCircumferenceCm, &biometrics.RestingHeartRateBpm, &biometrics.ActivityLevel, &biometrics.CreatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &biometrics, nil
}
