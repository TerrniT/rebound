package usecase

import (
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
)

type nutritionUseCase struct {
	repo entity.NutritionRepository
}

// CalculateDailyNutrition implements entity.NutritionUseCase.
func (uc *nutritionUseCase) CalculateDailyNutrition(userID int64, date time.Time) (*entity.Nutrition, error) {
	// TODO: Implement daily nutrition calculation
	panic("unimplemented")
}

func NewNutritionUseCase(repo entity.NutritionRepository) entity.NutritionUseCase {
	return &nutritionUseCase{repo: repo}
}

func (uc *nutritionUseCase) CreateNutrition(nutrition *entity.Nutrition) error {
	return uc.repo.Create(nutrition)
}

func (uc *nutritionUseCase) GetNutrition(id int64) (*entity.Nutrition, error) {
	return uc.repo.GetByID(id)
}

func (uc *nutritionUseCase) UpdateNutrition(nutrition *entity.Nutrition) error {
	return uc.repo.Update(nutrition)
}

func (uc *nutritionUseCase) DeleteNutrition(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *nutritionUseCase) ListNutritions(offset, limit int) ([]*entity.Nutrition, error) {
	return uc.repo.List(offset, limit)
}
