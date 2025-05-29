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