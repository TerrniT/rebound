package entity

import "time"

// MealFoodItem represents a food item in a meal
type MealFoodItem struct {
	ID                  string    `json:"id"`
	MealID              string    `json:"meal_id"`
	FoodItemID          string    `json:"food_item_id"`
	QuantityConsumed    float64   `json:"quantity_consumed"`
	ServingUnitConsumed string    `json:"serving_unit_consumed"`
	CaloriesConsumed    float64   `json:"calories_consumed"`
	ProteinConsumed     float64   `json:"protein_consumed"`
	FatConsumed         float64   `json:"fat_consumed"`
	CarbsConsumed       float64   `json:"carbs_consumed"`
	LoggedAt            time.Time `json:"logged_at"`
}
