-- begin of the transaction

BEGIN;

-- Drop Indexes (Indexes are dropped automatically with the table, but explicitly listing them
-- can be part of some migration strategies or if you want to drop them independent of tables)
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_auth_tokens_user_id;
DROP INDEX IF EXISTS idx_auth_tokens_token_hash;
DROP INDEX IF EXISTS idx_exercises_name;
DROP INDEX IF EXISTS idx_exercises_muscle_group_primary;
DROP INDEX IF EXISTS idx_exercises_type;
DROP INDEX IF EXISTS idx_workout_plans_user_id;
DROP INDEX IF EXISTS idx_workout_plans_name;
DROP INDEX IF EXISTS idx_workout_plans_type;
DROP INDEX IF EXISTS idx_wpe_plan_id;
DROP INDEX IF EXISTS idx_wpe_exercise_id;
DROP INDEX IF EXISTS idx_uws_user_id;
DROP INDEX IF EXISTS idx_uws_plan_id;
DROP INDEX IF EXISTS idx_uws_scheduled_at;
DROP INDEX IF EXISTS idx_uws_status;
DROP INDEX IF EXISTS idx_uwsl_session_id;
DROP INDEX IF EXISTS idx_uwsl_exercise_id;
DROP INDEX IF EXISTS idx_food_items_name;
DROP INDEX IF EXISTS idx_food_items_brand_name;
DROP INDEX IF EXISTS idx_food_items_barcode_upc;
DROP INDEX IF EXISTS idx_serving_units_food_item_id;
DROP INDEX IF EXISTS idx_serving_units_name;
DROP INDEX IF EXISTS idx_user_meals_user_id;
DROP INDEX IF EXISTS idx_user_meals_date;
DROP INDEX IF EXISTS idx_user_meals_type;
DROP INDEX IF EXISTS idx_mfi_meal_id;
DROP INDEX IF EXISTS idx_mfi_food_item_id;
DROP INDEX IF EXISTS idx_user_nutrition_goals_user_id;
DROP INDEX IF EXISTS idx_user_nutrition_goals_date;
DROP INDEX IF EXISTS idx_user_nutrition_goals_is_active;
DROP INDEX IF EXISTS idx_user_biometrics_user_id;
DROP INDEX IF EXISTS idx_user_biometrics_log_date;

-- Drop the unique constraint if it was added
ALTER TABLE UserBiometrics DROP CONSTRAINT IF EXISTS uc_user_biometrics_user_date;

COMMIT;