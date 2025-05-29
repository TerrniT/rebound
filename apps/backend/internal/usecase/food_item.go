package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/repository"
)

type Config struct {
	DefaultPageSize int
	MaxPageSize     int
}

type FoodItemUseCase struct {
	repo   repository.FoodItemRepository
	config Config
}

// New creates a new instance of FoodItemUseCase
func NewFoodItemUseCase(r repository.FoodItemRepository, config Config) *FoodItemUseCase {
	return &FoodItemUseCase{
		repo:   r,
		config: config,
	}
}

func (uc *FoodItemUseCase) CreateFoodItem(ctx context.Context, foodItem *entity.FoodItem) (*entity.FoodItem, error) {
	foodItem.ID = uuid.New().String()
	foodItem.CreatedAt = time.Now()
	foodItem.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, foodItem)
}

func (uc *FoodItemUseCase) GetFoodItem(ctx context.Context, foodItemID string) (*entity.FoodItem, error) {
	return uc.repo.GetByID(ctx, foodItemID)
}

func (uc *FoodItemUseCase) ListFoodItems(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*entity.FoodItem, int64, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = uc.config.DefaultPageSize
	}
	if pageSize > uc.config.MaxPageSize {
		pageSize = uc.config.MaxPageSize
	}

	// Get total count
	total, err := uc.repo.Count(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	// Get food items
	items, err := uc.repo.List(ctx, filters, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (uc *FoodItemUseCase) UpdateFoodItem(ctx context.Context, foodItem *entity.FoodItem) error {
	foodItem.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, foodItem)
}

func (uc *FoodItemUseCase) DeleteFoodItem(ctx context.Context, foodItemID string) error {
	return uc.repo.Delete(ctx, foodItemID)
}

func (uc *FoodItemUseCase) SearchFoodItems(ctx context.Context, query string, page, pageSize int) ([]*entity.FoodItem, int64, error) {
	// Validate page size
	if pageSize <= 0 {
		pageSize = uc.config.DefaultPageSize
	}
	if pageSize > uc.config.MaxPageSize {
		pageSize = uc.config.MaxPageSize
	}

	// Get total count
	total, err := uc.repo.CountBySearch(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	// Search food items
	items, err := uc.repo.Search(ctx, query, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// func (uc *FoodItemUseCase) CreateServingUnit(ctx context.Context, servingUnit *entity.ServingUnit) (*entity.ServingUnit, error) {
// 	// Verify food item exists if provided
// 	if servingUnit.FoodItemID != nil {
// 		exists, err := uc.repo.Exists(ctx, *servingUnit.FoodItemID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if !exists {
// 			return nil, errors.New("food item not found")
// 		}
// 	}

// 	return uc.repo.CreateServingUnit(ctx, servingUnit)
// }

// func (uc *FoodItemUseCase) GetServingUnits(ctx context.Context, foodItemID string) ([]*entity.ServingUnit, error) {
// 	// Verify food item exists
// 	exists, err := uc.repo.Exists(ctx, foodItemID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !exists {
// 		return nil, errors.New("food item not found")
// 	}

// 	return uc.repo.ListServingUnits(ctx, foodItemID)
// }

// func (uc *FoodItemUseCase) UpdateServingUnit(ctx context.Context, servingUnit *entity.ServingUnit) error {
// 	// Verify food item exists if provided
// 	if servingUnit.FoodItemID != nil {
// 		exists, err := uc.repo.Exists(ctx, *servingUnit.FoodItemID)
// 		if err != nil {
// 			return err
// 		}
// 		if !exists {
// 			return errors.New("food item not found")
// 		}
// 	}

// 	return uc.repo.UpdateServingUnit(ctx, servingUnit)
// }

// func (uc *FoodItemUseCase) DeleteServingUnit(ctx context.Context, servingUnitID int) error {
// 	return uc.repo.DeleteServingUnit(ctx, servingUnitID)
// }

// func (uc *FoodItemUseCase) CalculateNutrition(ctx context.Context, foodItemID string, quantity float64, servingUnit string) (*entity.FoodItem, error) {
// 	// Get food item
// 	foodItem, err := uc.repo.GetByID(ctx, foodItemID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get serving unit
// 	servingUnits, err := uc.repo.ListServingUnits(ctx, foodItemID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Find matching serving unit
// 	var targetUnit *entity.ServingUnit
// 	for _, unit := range servingUnits {
// 		if unit.UnitName == servingUnit || unit.Abbreviation == servingUnit {
// 			targetUnit = unit
// 			break
// 		}
// 	}

// 	if targetUnit == nil {
// 		return nil, errors.New("serving unit not found")
// 	}

// 	// Calculate conversion factor
// 	var conversionFactor float64
// 	if targetUnit.GramsEquivalent != nil {
// 		conversionFactor = *targetUnit.GramsEquivalent / foodItem.ServingSizeDefaultQty
// 	} else if targetUnit.MlEquivalent != nil {
// 		conversionFactor = *targetUnit.MlEquivalent / foodItem.ServingSizeDefaultQty
// 	} else {
// 		return nil, errors.New("serving unit has no conversion factor")
// 	}

// 	// Calculate nutritional values
// 	adjustedQuantity := quantity * conversionFactor
// 	foodItem.CaloriesPerDefaultServing *= adjustedQuantity
// 	foodItem.ProteinGramsPerDefaultServing *= adjustedQuantity
// 	foodItem.FatGramsPerDefaultServing *= adjustedQuantity
// 	foodItem.CarbsGramsPerDefaultServing *= adjustedQuantity

// 	if foodItem.FiberGramsPerDefaultServing != nil {
// 		*foodItem.FiberGramsPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.SugarGramsPerDefaultServing != nil {
// 		*foodItem.SugarGramsPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.SaturatedFatGramsPerDefaultServing != nil {
// 		*foodItem.SaturatedFatGramsPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.TransFatGramsPerDefaultServing != nil {
// 		*foodItem.TransFatGramsPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.CholesterolMgPerDefaultServing != nil {
// 		*foodItem.CholesterolMgPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.SodiumMgPerDefaultServing != nil {
// 		*foodItem.SodiumMgPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.PotassiumMgPerDefaultServing != nil {
// 		*foodItem.PotassiumMgPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.VitaminAMcgPerDefaultServing != nil {
// 		*foodItem.VitaminAMcgPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.VitaminCMgPerDefaultServing != nil {
// 		*foodItem.VitaminCMgPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.CalciumMgPerDefaultServing != nil {
// 		*foodItem.CalciumMgPerDefaultServing *= adjustedQuantity
// 	}
// 	if foodItem.IronMgPerDefaultServing != nil {
// 		*foodItem.IronMgPerDefaultServing *= adjustedQuantity
// 	}

// 	return foodItem, nil
// }
