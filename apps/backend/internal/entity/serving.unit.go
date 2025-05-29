package entity

// ServingUnit represents a serving unit for food items
type ServingUnit struct {
	ID              int      `json:"id"`
	FoodItemID      *string  `json:"food_item_id,omitempty"`
	UnitName        string   `json:"unit_name"`
	Abbreviation    string   `json:"abbreviation"`
	GramsEquivalent *float64 `json:"grams_equivalent,omitempty"`
	MlEquivalent    *float64 `json:"ml_equivalent,omitempty"`
}
