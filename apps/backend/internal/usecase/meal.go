package usecase

import (
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
)

type mealUseCase struct {
	repo          entity.MealRepository
	nutritionRepo entity.NutritionRepository
}

func NewMealUseCase(repo entity.MealRepository, nutritionRepo entity.NutritionRepository) entity.MealUseCase {
	return &mealUseCase{
		repo:          repo,
		nutritionRepo: nutritionRepo,
	}
}

func (uc *mealUseCase) CreateMeal(meal *entity.Meal) error {
	return uc.repo.Create(meal)
}

func (uc *mealUseCase) GetMeal(id int64) (*entity.Meal, error) {
	return uc.repo.GetByID(id)
}

func (uc *mealUseCase) UpdateMeal(meal *entity.Meal) error {
	return uc.repo.Update(meal)
}

func (uc *mealUseCase) DeleteMeal(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *mealUseCase) ListMeals(userID int64, startDate, endDate time.Time) ([]*entity.Meal, error) {
	return uc.repo.List(userID, startDate, endDate)
}

func (uc *mealUseCase) GetDailyMeals(userID int64, date time.Time) ([]*entity.Meal, error) {
	return uc.repo.GetByDate(userID, date)
}

func (uc *mealUseCase) CalculateDailyNutrition(userID int64, date time.Time) (*entity.Nutrition, error) {
	// TODO: Implement daily nutrition calculation
	return nil, nil
}
