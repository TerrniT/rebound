package entity

import "time"

// UserNutritionGoal represents a user's nutrition goals
type UserNutritionGoal struct {
	ID                    string    `json:"id"`
	UserID                string    `json:"user_id"`
	GoalEffectiveDate     time.Time `json:"goal_effective_date"`
	TargetCalories        float64   `json:"target_calories"`
	TargetProteinGrams    float64   `json:"target_protein_grams"`
	TargetFatGrams        float64   `json:"target_fat_grams"`
	TargetCarbsGrams      float64   `json:"target_carbs_grams"`
	TargetFiberGrams      *float64  `json:"target_fiber_grams,omitempty"`
	TargetSugarGramsLimit *float64  `json:"target_sugar_grams_limit,omitempty"`
	Notes                 *string   `json:"notes,omitempty"`
	IsActive              bool      `json:"is_active"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
