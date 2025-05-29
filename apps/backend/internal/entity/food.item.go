package entity

import "time"

// FoodItemSource represents the source of a food item
type FoodItemSource string

const (
	FoodItemSourceUserCreated FoodItemSource = "user_created"
	FoodItemSourceSystem      FoodItemSource = "system"
	FoodItemSourceAPI         FoodItemSource = "api"
)

// FoodItem represents a food item in the system
type FoodItem struct {
	ID                                 string         `json:"id"`
	Name                               string         `json:"name"`
	BrandName                          *string        `json:"brand_name,omitempty"`
	BarcodeUPC                         *string        `json:"barcode_upc,omitempty"`
	ServingSizeDefaultQty              float64        `json:"serving_size_default_qty"`
	ServingSizeDefaultUnit             string         `json:"serving_size_default_unit"`
	CaloriesPerDefaultServing          float64        `json:"calories_per_default_serving"`
	ProteinGramsPerDefaultServing      float64        `json:"protein_grams_per_default_serving"`
	FatGramsPerDefaultServing          float64        `json:"fat_grams_per_default_serving"`
	CarbsGramsPerDefaultServing        float64        `json:"carbs_grams_per_default_serving"`
	FiberGramsPerDefaultServing        *float64       `json:"fiber_grams_per_default_serving,omitempty"`
	SugarGramsPerDefaultServing        *float64       `json:"sugar_grams_per_default_serving,omitempty"`
	SaturatedFatGramsPerDefaultServing *float64       `json:"saturated_fat_grams_per_default_serving,omitempty"`
	TransFatGramsPerDefaultServing     *float64       `json:"trans_fat_grams_per_default_serving,omitempty"`
	CholesterolMgPerDefaultServing     *float64       `json:"cholesterol_mg_per_default_serving,omitempty"`
	SodiumMgPerDefaultServing          *float64       `json:"sodium_mg_per_default_serving,omitempty"`
	PotassiumMgPerDefaultServing       *float64       `json:"potassium_mg_per_default_serving,omitempty"`
	VitaminAMcgPerDefaultServing       *float64       `json:"vitamin_a_mcg_per_default_serving,omitempty"`
	VitaminCMgPerDefaultServing        *float64       `json:"vitamin_c_mg_per_default_serving,omitempty"`
	CalciumMgPerDefaultServing         *float64       `json:"calcium_mg_per_default_serving,omitempty"`
	IronMgPerDefaultServing            *float64       `json:"iron_mg_per_default_serving,omitempty"`
	Source                             FoodItemSource `json:"source"`
	IsVerified                         bool           `json:"is_verified"`
	CreatedByUserID                    *string        `json:"created_by_user_id,omitempty"`
	CreatedAt                          time.Time      `json:"created_at"`
	UpdatedAt                          time.Time      `json:"updated_at"`
}
