package entity

// DayOfWeek represents a day of the week
type DayOfWeek string

const (
	DayOfWeekMonday    DayOfWeek = "monday"
	DayOfWeekTuesday   DayOfWeek = "tuesday"
	DayOfWeekWednesday DayOfWeek = "wednesday"
	DayOfWeekThursday  DayOfWeek = "thursday"
	DayOfWeekFriday    DayOfWeek = "friday"
	DayOfWeekSaturday  DayOfWeek = "saturday"
	DayOfWeekSunday    DayOfWeek = "sunday"
)

// WorkoutPlanExercise represents an exercise in a workout plan
type WorkoutPlanExercise struct {
	ID                string     `json:"id"`
	PlanID            string     `json:"plan_id"`
	ExerciseID        string     `json:"exercise_id"`
	DayOfWeek         *DayOfWeek `json:"day_of_week,omitempty"`
	DayNumber         *int       `json:"day_number,omitempty"`
	ExerciseOrder     int        `json:"exercise_order"`
	Sets              *int       `json:"sets,omitempty"`
	RepsMin           *int       `json:"reps_min,omitempty"`
	RepsMax           *int       `json:"reps_max,omitempty"`
	RepsTarget        *int       `json:"reps_target,omitempty"`
	DurationSeconds   *int       `json:"duration_seconds,omitempty"`
	RestPeriodSeconds *int       `json:"rest_period_seconds,omitempty"`
	Notes             *string    `json:"notes,omitempty"`
}
