package entity

import "time"

// WorkoutPlan represents a workout plan
type WorkoutPlan struct {
	ID               string             `json:"id"`
	UserID           *string            `json:"user_id,omitempty"`
	Name             string             `json:"name"`
	Description      *string            `json:"description,omitempty"`
	Type             *string            `json:"type,omitempty"`
	DifficultyLevel  ExerciseDifficulty `json:"difficulty_level"`
	DurationEstimate *int               `json:"duration_estimate_minutes,omitempty"`
	FrequencyPerWeek *int               `json:"frequency_per_week,omitempty"`
	IsPublic         bool               `json:"is_public"`
	CoverImageURL    *string            `json:"cover_image_url,omitempty"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}
