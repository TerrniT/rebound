package entity

import "time"

// ActivityLevel represents the activity level of a user
type ActivityLevel string

const (
	ActivityLevelSedentary        ActivityLevel = "sedentary"
	ActivityLevelLightlyActive    ActivityLevel = "lightly_active"
	ActivityLevelModeratelyActive ActivityLevel = "moderately_active"
	ActivityLevelVeryActive       ActivityLevel = "very_active"
	ActivityLevelExtraActive      ActivityLevel = "extra_active"
)

// UserBiometric represents a user's biometric data
type UserBiometric struct {
	ID                   string         `json:"id"`
	UserID               string         `json:"user_id"`
	LogDate              time.Time      `json:"log_date"`
	WeightKg             *float64       `json:"weight_kg,omitempty"`
	HeightCm             *float64       `json:"height_cm,omitempty"`
	BodyFatPercentage    *float64       `json:"body_fat_percentage,omitempty"`
	WaistCircumferenceCm *float64       `json:"waist_circumference_cm,omitempty"`
	HipCircumferenceCm   *float64       `json:"hip_circumference_cm,omitempty"`
	ChestCircumferenceCm *float64       `json:"chest_circumference_cm,omitempty"`
	RestingHeartRateBpm  *int           `json:"resting_heart_rate_bpm,omitempty"`
	ActivityLevel        *ActivityLevel `json:"activity_level,omitempty"`
	CreatedAt            time.Time      `json:"created_at"`
}
