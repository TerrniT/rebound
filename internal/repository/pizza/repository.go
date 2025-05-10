package pizza

import (
	"context"

	"github.com/terrnit/pizza-crud-go/internal/domain/pizza"
)

// Repository defines the interface for pizza data operations
type Repository interface {
	Create(ctx context.Context, pizza *pizza.Pizza) error
	GetByID(ctx context.Context, id int64) (*pizza.Pizza, error)
	Update(ctx context.Context, pizza *pizza.Pizza) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*pizza.Pizza, error)
}
