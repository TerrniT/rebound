package entity

import "time"

// UserWorkoutSessionLog represents a log entry for a workout session
type UserWorkoutSessionLog struct {
	ID                       string    `json:"id"`
	SessionID                string    `json:"session_id"`
	ExerciseID               string    `json:"exercise_id"`
	PlanExerciseID           *string   `json:"plan_exercise_id,omitempty"`
	SetNumber                int       `json:"set_number"`
	RepsCompleted            *int      `json:"reps_completed,omitempty"`
	WeightKg                 *float64  `json:"weight_kg,omitempty"`
	DistanceKm               *float64  `json:"distance_km,omitempty"`
	DurationSecondsCompleted *int      `json:"duration_seconds_completed,omitempty"`
	RestTakenSeconds         *int      `json:"rest_taken_seconds,omitempty"`
	Notes                    *string   `json:"notes,omitempty"`
	LoggedAt                 time.Time `json:"logged_at"`
}
