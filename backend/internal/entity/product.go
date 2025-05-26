package entity

import "time"

// Product represents a food product in the system
type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Nutrition   Nutrition `json:"nutrition"`
	Barcode     string    `json:"barcode,omitempty"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductRepository interface defines the methods for product data persistence
type ProductRepository interface {
	Create(product *Product) error
	GetByID(id int64) (*Product, error)
	Update(product *Product) error
	Delete(id int64) error
	List(offset, limit int) ([]*Product, error)
	Search(query string) ([]*Product, error)
	GetByBarcode(barcode string) (*Product, error)
}

// ProductUseCase interface defines the business logic for products
type ProductUseCase interface {
	CreateProduct(product *Product) error
	GetProduct(id int64) (*Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id int64) error
	ListProducts(offset, limit int) ([]*Product, error)
	SearchProducts(query string) ([]*Product, error)
	GetProductByBarcode(barcode string) (*Product, error)
}
