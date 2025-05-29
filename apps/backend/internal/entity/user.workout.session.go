package entity

import "time"

// WorkoutSessionStatus represents the status of a workout session
type WorkoutSessionStatus string

const (
	WorkoutSessionStatusScheduled  WorkoutSessionStatus = "scheduled"
	WorkoutSessionStatusInProgress WorkoutSessionStatus = "in_progress"
	WorkoutSessionStatusCompleted  WorkoutSessionStatus = "completed"
	WorkoutSessionStatusCancelled  WorkoutSessionStatus = "cancelled"
)

// UserWorkoutSession represents a user's workout session
type UserWorkoutSession struct {
	ID                      string               `json:"id"`
	UserID                  string               `json:"user_id"`
	PlanID                  *string              `json:"plan_id,omitempty"`
	SessionName             *string              `json:"session_name,omitempty"`
	ScheduledAt             *time.Time           `json:"scheduled_at,omitempty"`
	StartedAt               *time.Time           `json:"started_at,omitempty"`
	CompletedAt             *time.Time           `json:"completed_at,omitempty"`
	DurationMinutes         *int                 `json:"duration_minutes,omitempty"`
	Status                  WorkoutSessionStatus `json:"status"`
	Notes                   *string              `json:"notes,omitempty"`
	Location                *string              `json:"location,omitempty"`
	MoodRating              *int                 `json:"mood_rating,omitempty"`
	PerceivedExertionRating *int                 `json:"perceived_exertion_rating,omitempty"`
	CreatedAt               time.Time            `json:"created_at"`
	UpdatedAt               time.Time            `json:"updated_at"`
}
