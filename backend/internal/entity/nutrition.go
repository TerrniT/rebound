package entity

import "time"

// Nutrition represents the nutritional information for a food item
type Nutrition struct {
	ID          int64     `json:"id"`
	Calories    float64   `json:"calories"`
	Proteins    float64   `json:"proteins"`
	Carbs       float64   `json:"carbs"`
	Fats        float64   `json:"fats"`
	ServingSize float64   `json:"serving_size"`
	Unit        string    `json:"unit"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NutritionRepository interface defines the methods for nutrition data persistence
type NutritionRepository interface {
	Create(nutrition *Nutrition) error
	GetByID(id int64) (*Nutrition, error)
	Update(nutrition *Nutrition) error
	Delete(id int64) error
	List(offset, limit int) ([]*Nutrition, error)
}

// NutritionUseCase interface defines the business logic for nutrition
type NutritionUseCase interface {
	CreateNutrition(nutrition *Nutrition) error
	GetNutrition(id int64) (*Nutrition, error)
	UpdateNutrition(nutrition *Nutrition) error
	DeleteNutrition(id int64) error
	ListNutritions(offset, limit int) ([]*Nutrition, error)
	CalculateDailyNutrition(userID int64, date time.Time) (*Nutrition, error)
}
