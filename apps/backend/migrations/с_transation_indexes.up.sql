-- begin of the transaction

BEGIN;

-- Add additional Indexes for query performance
CREATE INDEX idx_users_username ON Users(username);
CREATE INDEX idx_users_email ON Users(email);
CREATE INDEX idx_auth_tokens_user_id ON AuthTokens(user_id);
CREATE INDEX idx_auth_tokens_token_hash ON AuthTokens(token_hash); -- Unique index also helps here

CREATE INDEX idx_exercises_name ON Exercises(exercise_name);
CREATE INDEX idx_exercises_muscle_group_primary ON Exercises(muscle_group_primary);
CREATE INDEX idx_exercises_type ON Exercises(exercise_type);
CREATE INDEX idx_workout_plans_user_id ON WorkoutPlans(user_id);
CREATE INDEX idx_workout_plans_name ON WorkoutPlans(plan_name);
CREATE INDEX idx_workout_plans_type ON WorkoutPlans(plan_type);
CREATE INDEX idx_wpe_plan_id ON WorkoutPlanExercises(plan_id);
CREATE INDEX idx_wpe_exercise_id ON WorkoutPlanExercises(exercise_id);
CREATE INDEX idx_uws_user_id ON UserWorkoutSessions(user_id);
CREATE INDEX idx_uws_plan_id ON UserWorkoutSessions(plan_id);
CREATE INDEX idx_uws_scheduled_at ON UserWorkoutSessions(scheduled_at);
CREATE INDEX idx_uws_status ON UserWorkoutSessions(status);
CREATE INDEX idx_uwsl_session_id ON UserWorkoutSessionLogs(session_id);
CREATE INDEX idx_uwsl_exercise_id ON UserWorkoutSessionLogs(exercise_id);

CREATE INDEX idx_food_items_name ON FoodItems(food_name);
CREATE INDEX idx_food_items_brand_name ON FoodItems(brand_name);
CREATE INDEX idx_food_items_barcode_upc ON FoodItems(barcode_upc);
CREATE INDEX idx_serving_units_food_item_id ON ServingUnits(food_item_id);
CREATE INDEX idx_serving_units_name ON ServingUnits(unit_name);
CREATE INDEX idx_user_meals_user_id ON UserMeals(user_id);
CREATE INDEX idx_user_meals_date ON UserMeals(meal_date);
CREATE INDEX idx_user_meals_type ON UserMeals(meal_type);
CREATE INDEX idx_mfi_meal_id ON MealFoodItems(meal_id);
CREATE INDEX idx_mfi_food_item_id ON MealFoodItems(food_item_id);
CREATE INDEX idx_user_nutrition_goals_user_id ON UserNutritionGoals(user_id);
CREATE INDEX idx_user_nutrition_goals_date ON UserNutritionGoals(goal_effective_date);
CREATE INDEX idx_user_nutrition_goals_is_active ON UserNutritionGoals(is_active);
CREATE INDEX idx_user_biometrics_user_id ON UserBiometrics(user_id);
CREATE INDEX idx_user_biometrics_log_date ON UserBiometrics(log_date);

-- Add a unique constraint for user biometrics log per day if needed
ALTER TABLE UserBiometrics ADD CONSTRAINT uc_user_biometrics_user_date UNIQUE (user_id, log_date);


COMMIT;