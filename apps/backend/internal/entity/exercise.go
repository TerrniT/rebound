package entity

import "time"

// ExerciseDifficulty represents the difficulty level of an exercise
type ExerciseDifficulty string

const (
	ExerciseDifficultyBeginner     ExerciseDifficulty = "beginner"
	ExerciseDifficultyIntermediate ExerciseDifficulty = "intermediate"
	ExerciseDifficultyAdvanced     ExerciseDifficulty = "advanced"
	ExerciseDifficultyExpert       ExerciseDifficulty = "expert"
)

// ExerciseType represents the type of exercise
type ExerciseType string

const (
	ExerciseTypeStrength    ExerciseType = "strength"
	ExerciseTypeCardio      ExerciseType = "cardio"
	ExerciseTypeFlexibility ExerciseType = "flexibility"
	ExerciseTypeBalance     ExerciseType = "balance"
	ExerciseTypeHIIT        ExerciseType = "hiit"
	ExerciseTypeYoga        ExerciseType = "yoga"
	ExerciseTypePilates     ExerciseType = "pilates"
	ExerciseTypeOther       ExerciseType = "other"
)

// Exercise represents a physical exercise
type Exercise struct {
	ID                    string             `json:"id"`
	Name                  string             `json:"name"`
	Description           *string            `json:"description,omitempty"`
	MuscleGroupPrimary    *string            `json:"muscle_group_primary,omitempty"`
	MuscleGroupsSecondary []string           `json:"muscle_groups_secondary,omitempty"`
	EquipmentRequired     string             `json:"equipment_required"`
	DifficultyLevel       ExerciseDifficulty `json:"difficulty_level"`
	VideoURL              *string            `json:"video_url,omitempty"`
	ImageURLThumbnail     *string            `json:"image_url_thumbnail,omitempty"`
	ImageURLMain          *string            `json:"image_url_main,omitempty"`
	Type                  ExerciseType       `json:"type"`
	CreatedByUserID       *string            `json:"created_by_user_id,omitempty"`
	IsPublic              bool               `json:"is_public"`
	CreatedAt             time.Time          `json:"created_at"`
	UpdatedAt             time.Time          `json:"updated_at"`
}
