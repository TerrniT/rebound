package pizza

import "time"

// Pizza represents the pizza entity
type Pizza struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewPizza creates a new pizza instance
func NewPizza(name, description string, price float64) *Pizza {
	now := time.Now()
	return &Pizza{
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
