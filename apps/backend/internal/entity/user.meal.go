package entity

import "time"

// UserMealType represents the type of a user's meal
type UserMealType string

const (
	UserMealTypeBreakfast UserMealType = "breakfast"
	UserMealTypeLunch     UserMealType = "lunch"
	UserMealTypeDinner    UserMealType = "dinner"
	UserMealTypeSnack     UserMealType = "snack"
	UserMealTypeOther     UserMealType = "other"
)

// UserMeal represents a user's meal
type UserMeal struct {
	ID                    string       `json:"id"`
	UserID                string       `json:"user_id"`
	MealType              UserMealType `json:"meal_type"`
	MealDate              time.Time    `json:"meal_date"`
	MealTime              *time.Time   `json:"meal_time,omitempty"`
	CustomMealName        *string      `json:"custom_meal_name,omitempty"`
	Notes                 *string      `json:"notes,omitempty"`
	TotalCaloriesConsumed *float64     `json:"total_calories_consumed,omitempty"`
	TotalProteinConsumed  *float64     `json:"total_protein_consumed,omitempty"`
	TotalFatConsumed      *float64     `json:"total_fat_consumed,omitempty"`
	TotalCarbsConsumed    *float64     `json:"total_carbs_consumed,omitempty"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
}
