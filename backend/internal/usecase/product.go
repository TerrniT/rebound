package usecase

import (
	"github.com/terrnit/rebound/backend/internal/entity"
)

type productUseCase struct {
	repo          entity.ProductRepository
	nutritionRepo entity.NutritionRepository
}

func NewProductUseCase(repo entity.ProductRepository, nutritionRepo entity.NutritionRepository) entity.ProductUseCase {
	return &productUseCase{
		repo:          repo,
		nutritionRepo: nutritionRepo,
	}
}

func (uc *productUseCase) CreateProduct(product *entity.Product) error {
	return uc.repo.Create(product)
}

func (uc *productUseCase) GetProduct(id int64) (*entity.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *productUseCase) UpdateProduct(product *entity.Product) error {
	return uc.repo.Update(product)
}

func (uc *productUseCase) DeleteProduct(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *productUseCase) ListProducts(offset, limit int) ([]*entity.Product, error) {
	return uc.repo.List(offset, limit)
}

func (uc *productUseCase) SearchProducts(query string) ([]*entity.Product, error) {
	return uc.repo.Search(query)
}

func (uc *productUseCase) GetProductByBarcode(barcode string) (*entity.Product, error) {
	return uc.repo.GetByBarcode(barcode)
}
