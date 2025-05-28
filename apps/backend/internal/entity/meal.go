package entity

import "time"

// MealType represents different types of meals (breakfast, lunch, dinner, snack)
type MealType string

const (
	Breakfast MealType = "breakfast"
	Lunch     MealType = "lunch"
	Dinner    MealType = "dinner"
	Snack     MealType = "snack"
)

// Meal represents a meal entry in the system
type Meal struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Type        MealType  `json:"type"`
	Nutrition   Nutrition `json:"nutrition"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MealRepository interface defines the methods for meal data persistence
type MealRepository interface {
	Create(meal *Meal) error
	GetByID(id int64) (*Meal, error)
	Update(meal *Meal) error
	Delete(id int64) error
	List(userID int64, startDate, endDate time.Time) ([]*Meal, error)
	GetByDate(userID int64, date time.Time) ([]*Meal, error)
}

// MealUseCase interface defines the business logic for meals
type MealUseCase interface {
	CreateMeal(meal *Meal) error
	GetMeal(id int64) (*Meal, error)
	UpdateMeal(meal *Meal) error
	DeleteMeal(id int64) error
	ListMeals(userID int64, startDate, endDate time.Time) ([]*Meal, error)
	GetDailyMeals(userID int64, date time.Time) ([]*Meal, error)
	CalculateDailyNutrition(userID int64, date time.Time) (*Nutrition, error)
}
