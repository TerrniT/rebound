-- Down Migration file to drop the fitness app database schema for PostgreSQL

-- Use a transaction for atomicity
BEGIN;

-- Drop Tables (must be done in reverse order of dependencies or use CASCADE)
-- Dropping child tables first is generally safer and more explicit
-- Dependencies:
-- UserWorkoutSessionLogs depends on UserWorkoutSessions, Exercises, WorkoutPlanExercises
-- MealFoodItems depends on UserMeals, FoodItems
-- WorkoutPlanExercises depends on WorkoutPlans, Exercises
-- UserWorkoutSessions depends on Users, WorkoutPlans
-- UserMeals depends on Users
-- UserNutritionGoals depends on Users
-- UserBiometrics depends on Users
-- ServingUnits depends on FoodItems
-- AuthTokens depends on Users
-- UserRoles depends on Users, Roles
-- WorkoutPlans depends on Users
-- Exercises depends on Users
-- FoodItems depends on Users
-- Roles has no outgoing FKs
-- Users has outgoing FKs to many, but nothing depends *only* on Users

-- Drop tables in an order that respects foreign key dependencies
DROP TABLE IF EXISTS UserWorkoutSessionLogs CASCADE; -- CASCADE will drop FKs pointing to this table (none)
DROP TABLE IF EXISTS MealFoodItems CASCADE;
DROP TABLE IF EXISTS WorkoutPlanExercises CASCADE;
DROP TABLE IF EXISTS UserNutritionGoals CASCADE;
DROP TABLE IF EXISTS UserBiometrics CASCADE;
DROP TABLE IF EXISTS ServingUnits CASCADE;
DROP TABLE IF EXISTS AuthTokens CASCADE;
DROP TABLE IF EXISTS UserRoles CASCADE;
DROP TABLE IF EXISTS UserWorkoutSessions CASCADE; -- CASCADE will drop FKs from logs table
DROP TABLE IF EXISTS UserMeals CASCADE;         -- CASCADE will drop FKs from MealFoodItems table
DROP TABLE IF EXISTS WorkoutPlans CASCADE;       -- CASCADE will drop FKs from WorkoutPlanExercises, UserWorkoutSessions
DROP TABLE IF EXISTS Exercises CASCADE;          -- CASCADE will drop FKs from WorkoutPlanExercises, UserWorkoutSessionLogs, MealFoodItems (indirectly via plan_exercise_id if used)
DROP TABLE IF EXISTS FoodItems CASCADE;          -- CASCADE will drop FKs from MealFoodItems, ServingUnits
DROP TABLE IF EXISTS Roles CASCADE;
DROP TABLE IF EXISTS Users CASCADE;              -- CASCADE will drop FKs from all user-related tables

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


-- Drop ENUM types (must be done after tables that use them are dropped)
DROP TYPE IF EXISTS user_gender_enum;
DROP TYPE IF EXISTS auth_token_type_enum;
DROP TYPE IF EXISTS exercise_difficulty_enum;
DROP TYPE IF EXISTS exercise_type_enum;
DROP TYPE IF EXISTS workout_session_status_enum;
DROP TYPE IF EXISTS food_item_source_enum;
DROP TYPE IF EXISTS user_meal_type_enum;
DROP TYPE IF EXISTS activity_level_enum;
DROP TYPE IF EXISTS day_of_week_enum;


-- Commit the transaction
COMMIT;