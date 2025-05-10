package pizza

import (
	"context"
	"sync"

	"github.com/terrnit/pizza-crud-go/internal/domain/pizza"
)

type memoryRepository struct {
	mu     sync.RWMutex
	pizzas map[int64]*pizza.Pizza
	nextID int64
}

// NewMemoryRepository creates a new in-memory pizza repository
func NewMemoryRepository() Repository {
	return &memoryRepository{
		pizzas: make(map[int64]*pizza.Pizza),
		nextID: 1,
	}
}

func (r *memoryRepository) Create(ctx context.Context, p *pizza.Pizza) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	p.ID = r.nextID
	r.nextID++
	r.pizzas[p.ID] = p
	return nil
}

func (r *memoryRepository) GetByID(ctx context.Context, id int64) (*pizza.Pizza, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if p, exists := r.pizzas[id]; exists {
		return p, nil
	}
	return nil, nil
}

func (r *memoryRepository) Update(ctx context.Context, p *pizza.Pizza) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.pizzas[p.ID]; !exists {
		return nil
	}
	r.pizzas[p.ID] = p
	return nil
}

func (r *memoryRepository) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.pizzas, id)
	return nil
}

func (r *memoryRepository) List(ctx context.Context) ([]*pizza.Pizza, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	pizzas := make([]*pizza.Pizza, 0, len(r.pizzas))
	for _, p := range r.pizzas {
		pizzas = append(pizzas, p)
	}
	return pizzas, nil
}
