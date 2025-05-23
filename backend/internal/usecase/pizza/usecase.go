package usecase

import (
	"context"
	"errors"

	"github.com/terrnit/rebound/backend/internal/domain/pizza"
	pizzarepo "github.com/terrnit/rebound/backend/internal/repository/pizza"
)

var (
	ErrPizzaNotFound = errors.New("pizza not found")
	ErrInvalidInput  = errors.New("invalid input")
)

// UseCase implements the pizza use cases
type UseCase struct {
	repo pizzarepo.Repository
}

// NewUseCase creates a new pizza use case
func NewUseCase(repo pizzarepo.Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

// Create creates a new pizza
func (uc *UseCase) Create(ctx context.Context, name, description string, price float64) (*pizza.Pizza, error) {
	if name == "" || price <= 0 {
		return nil, ErrInvalidInput
	}

	p := pizza.NewPizza(name, description, price)
	if err := uc.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

// GetByID retrieves a pizza by its ID
func (uc *UseCase) GetByID(ctx context.Context, id int64) (*pizza.Pizza, error) {
	p, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, ErrPizzaNotFound
	}
	return p, nil
}

// Update updates an existing pizza
func (uc *UseCase) Update(ctx context.Context, p *pizza.Pizza) error {
	if p.Name == "" || p.Price <= 0 {
		return ErrInvalidInput
	}
	return uc.repo.Update(ctx, p)
}

// Delete deletes a pizza by its ID
func (uc *UseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

// List retrieves all pizzas
func (uc *UseCase) List(ctx context.Context) ([]*pizza.Pizza, error) {
	return uc.repo.List(ctx)
}
