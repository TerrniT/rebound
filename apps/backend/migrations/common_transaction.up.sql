-- Migration file to set up the fitness app database schema for PostgreSQL

-- Use a transaction for atomicity
BEGIN;

-- Create ENUM types for columns with limited discrete values
-- These improve data integrity and readability
CREATE TYPE user_gender_enum AS ENUM ('Male', 'Female');
CREATE TYPE auth_token_type_enum AS ENUM ('access', 'refresh', 'password_reset', 'email_verification');
CREATE TYPE exercise_difficulty_enum AS ENUM ('Beginner', 'Intermediate', 'Advanced', 'All');
CREATE TYPE exercise_type_enum AS ENUM ('Strength', 'Cardio', 'Stretching', 'Plyometrics', 'Other');
CREATE TYPE workout_session_status_enum AS ENUM ('Scheduled', 'In Progress', 'Completed', 'Skipped', 'Cancelled');
CREATE TYPE food_item_source_enum AS ENUM ('USDA', 'OpenFoodFacts', 'UserAdded', 'Admin', 'Other');
CREATE TYPE user_meal_type_enum AS ENUM ('Breakfast', 'Lunch', 'Dinner', 'Snack', 'Pre-workout', 'Post-workout', 'Other');
CREATE TYPE activity_level_enum AS ENUM ('Sedentary', 'Lightly Active', 'Moderately Active', 'Very Active', 'Extremely Active');
CREATE TYPE day_of_week_enum AS ENUM ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday');


-- Create Tables

-- Auth Module
CREATE TABLE Users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    date_of_birth DATE,
    gender user_gender_enum,
    profile_picture_url VARCHAR(512),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    last_login_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE UserRoles (
    user_id UUID NOT NULL,
    role_id INT NOT NULL,
    assigned_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id) -- Composite primary key
);

CREATE TABLE AuthTokens (
    token_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token_type auth_token_type_enum NOT NULL,
    token_hash VARCHAR(255) NOT NULL UNIQUE, -- Store hash of the token
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    issued_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE
);


-- Trainings Module
CREATE TABLE Exercises (
    exercise_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    exercise_name VARCHAR(255) NOT NULL,
    description TEXT,
    muscle_group_primary VARCHAR(100),
    muscle_groups_secondary JSONB, -- Use JSONB for array/list of secondary muscles
    equipment_required VARCHAR(255) DEFAULT 'Bodyweight',
    difficulty_level exercise_difficulty_enum,
    video_url VARCHAR(512),
    image_url_thumbnail VARCHAR(512),
    image_url_main VARCHAR(512),
    exercise_type exercise_type_enum,
    created_by_user_id UUID, -- Nullable: system exercises exist
    is_public BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE WorkoutPlans (
    plan_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID, -- Nullable: system plans exist
    plan_name VARCHAR(255) NOT NULL,
    description TEXT,
    plan_type VARCHAR(100),
    difficulty_level exercise_difficulty_enum,
    duration_estimate_minutes INT,
    frequency_per_week INT,
    is_public BOOLEAN NOT NULL DEFAULT FALSE, -- User plans are private by default
    cover_image_url VARCHAR(512),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE WorkoutPlanExercises (
    -- Could be a composite PK (plan_id, exercise_id, exercise_order), but a surrogate key is often easier
    plan_exercise_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plan_id UUID NOT NULL,
    exercise_id UUID NOT NULL,
    day_of_week day_of_week_enum, -- Optional for weekly split plans
    day_number INT, -- Optional for multi-day split (e.g., 1, 2, 3)
    exercise_order INT NOT NULL DEFAULT 0, -- Order within the plan/day
    sets INT,
    reps_min INT,
    reps_max INT,
    reps_target INT,
    duration_seconds INT, -- For timed exercises
    rest_period_seconds INT,
    notes TEXT
    -- No created_at/updated_at for this join table, as its lifecycle is tied to the plan/exercise
);

CREATE TABLE UserWorkoutSessions (
    session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    plan_id UUID, -- Nullable: user can do un-planned workouts
    session_name VARCHAR(255), -- Custom name or derived from plan_name
    scheduled_at TIMESTAMP WITH TIME ZONE,
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    duration_minutes INT, -- Actual duration
    status workout_session_status_enum NOT NULL DEFAULT 'Scheduled',
    notes TEXT, -- User's general notes for the session
    location VARCHAR(255), -- e.g., Gym, Home
    mood_rating INT CHECK (mood_rating BETWEEN 1 AND 5),
    perceived_exertion_rating INT CHECK (perceived_exertion_rating BETWEEN 1 AND 10),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE UserWorkoutSessionLogs (
    -- Could be composite PK (session_id, exercise_id, set_number), but surrogate key is often easier
    log_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID NOT NULL,
    exercise_id UUID NOT NULL,
    plan_exercise_id UUID, -- Nullable: refers to the specific planned instance if from a plan
    set_number INT NOT NULL, -- Set sequence within the exercise for this session
    reps_completed INT,
    weight_kg DECIMAL(6,2),
    distance_km DECIMAL(7,3), -- For cardio
    duration_seconds_completed INT, -- For timed sets/exercises
    rest_taken_seconds INT, -- Actual rest taken after this set
    notes TEXT, -- Specific notes for this set
    logged_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP -- Time when this set was logged
);


-- Nutrition Module
CREATE TABLE FoodItems (
    food_item_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    food_name VARCHAR(255) NOT NULL,
    brand_name VARCHAR(150),
    barcode_upc VARCHAR(50) UNIQUE NULLS NOT DISTINCT, -- Allows multiple NULLs but only one of any specific value
    serving_size_default_qty DECIMAL(10,2) NOT NULL,
    serving_size_default_unit VARCHAR(50) NOT NULL,
    calories_per_default_serving DECIMAL(8,2) NOT NULL,
    protein_grams_per_default_serving DECIMAL(8,2) NOT NULL,
    fat_grams_per_default_serving DECIMAL(8,2) NOT NULL,
    carbs_grams_per_default_serving DECIMAL(8,2) NOT NULL,
    fiber_grams_per_default_serving DECIMAL(8,2),
    sugar_grams_per_default_serving DECIMAL(8,2),
    saturated_fat_grams_per_default_serving DECIMAL(8,2),
    trans_fat_grams_per_default_serving DECIMAL(8,2),
    cholesterol_mg_per_default_serving DECIMAL(8,2),
    sodium_mg_per_default_serving DECIMAL(8,2),
    potassium_mg_per_default_serving DECIMAL(8,2),
    vitamin_a_mcg_per_default_serving DECIMAL(10,2),
    vitamin_c_mg_per_default_serving DECIMAL(10,2),
    calcium_mg_per_default_serving DECIMAL(10,2),
    iron_mg_per_default_serving DECIMAL(10,2),
    source food_item_source_enum, -- e.g., USDA, OpenFoodFacts, UserAdded
    is_verified BOOLEAN NOT NULL DEFAULT FALSE, -- For user-added items
    created_by_user_id UUID, -- Nullable: for user-added items
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ServingUnits (
    unit_id SERIAL PRIMARY KEY,
    food_item_id UUID, -- Nullable: for generic units (g, ml, oz) vs food-specific (1 large apple)
    unit_name VARCHAR(50) NOT NULL,
    abbreviation VARCHAR(10) NOT NULL,
    grams_equivalent DECIMAL(10,4), -- Conversion factor to grams
    ml_equivalent DECIMAL(10,4) -- Conversion factor to milliliters
);

CREATE TABLE UserMeals (
    meal_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    meal_type user_meal_type_enum NOT NULL,
    meal_date DATE NOT NULL,
    meal_time TIME, -- Specific time of the meal
    custom_meal_name VARCHAR(150),
    notes TEXT,
    -- Total macros could be calculated via a view or trigger/function,
    -- or stored here and updated on MealFoodItems changes for read performance.
    -- Storing here is a common denormalization for reporting/display.
    total_calories_consumed DECIMAL(8,2),
    total_protein_consumed DECIMAL(8,2),
    total_fat_consumed DECIMAL(8,2),
    total_carbs_consumed DECIMAL(8,2),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE MealFoodItems (
    meal_food_item_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    meal_id UUID NOT NULL,
    food_item_id UUID NOT NULL,
    quantity_consumed DECIMAL(10,2) NOT NULL,
    serving_unit_consumed VARCHAR(50) NOT NULL, -- e.g., 'grams', 'cup', 'piece'. Should align with FoodItem/ServingUnit.
    -- Calculated macros based on quantity and food item data
    calories_consumed DECIMAL(8,2) NOT NULL,
    protein_consumed DECIMAL(8,2) NOT NULL,
    fat_consumed DECIMAL(8,2) NOT NULL,
    carbs_consumed DECIMAL(8,2) NOT NULL,
    logged_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE UserNutritionGoals (
    goal_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    goal_effective_date DATE NOT NULL DEFAULT CURRENT_DATE, -- Date from which this goal applies
    target_calories DECIMAL(8,2) NOT NULL,
    target_protein_grams DECIMAL(8,2) NOT NULL,
    target_fat_grams DECIMAL(8,2) NOT NULL,
    target_carbs_grams DECIMAL(8,2) NOT NULL,
    target_fiber_grams DECIMAL(8,2),
    target_sugar_grams_limit DECIMAL(8,2),
    notes TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE, -- Allows tracking historical goals
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
    -- Consider a UNIQUE constraint on (user_id, goal_effective_date) or manage 'is_active' uniqueness in application logic
);

CREATE TABLE UserBiometrics (
    biometric_log_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    log_date DATE NOT NULL DEFAULT CURRENT_DATE,
    weight_kg DECIMAL(6,2),
    height_cm DECIMAL(5,1), -- Height usually doesn't change often
    body_fat_percentage DECIMAL(4,2),
    waist_circumference_cm DECIMAL(5,1),
    hip_circumference_cm DECIMAL(5,1),
    chest_circumference_cm DECIMAL(5,1),
    resting_heart_rate_bpm INT,
    activity_level activity_level_enum, -- Used for TDEE calculation etc.
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
    -- Add a unique constraint if only one log per user per day is allowed
);


-- Add Foreign Key Constraints

-- Auth Module FKs
ALTER TABLE UserRoles ADD CONSTRAINT fk_user_roles_user_id FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE;
ALTER TABLE UserRoles ADD CONSTRAINT fk_user_roles_role_id FOREIGN KEY (role_id) REFERENCES Roles(role_id); -- ON DELETE RESTRICT/NO ACTION is typical for lookup tables
ALTER TABLE AuthTokens ADD CONSTRAINT fk_auth_tokens_user_id FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE;

-- Trainings Module FKs
ALTER TABLE Exercises ADD CONSTRAINT fk_exercises_created_by_user_id FOREIGN KEY (created_by_user_id) REFERENCES Users(user_id) ON DELETE SET NULL; -- If user deleted, keep exercise but remove creator link
ALTER TABLE WorkoutPlans ADD CONSTRAINT fk_workout_plans_user_id FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE SET NULL; -- If user deleted, keep plan but remove creator link (for shared plans)
ALTER TABLE WorkoutPlanExercises ADD CONSTRAINT fk_wpe_plan_id FOREIGN KEY (plan_id) REFERENCES WorkoutPlans(plan_id) ON DELETE CASCADE; -- If plan deleted, delete associated exercises in the plan
ALTER TABLE WorkoutPlanExercises ADD CONSTRAINT fk_wpe_exercise_id FOREIGN KEY (exercise_id) REFERENCES Exercises(exercise_id) ON DELETE RESTRICT; -- Prevent deleting an exercise if it's part of a plan
ALTER TABLE UserWorkoutSessions ADD CONSTRAINT fk_uws_user_id FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE; -- If user deleted, delete their session logs
ALTER TABLE UserWorkoutSessions ADD CONSTRAINT fk_uws_plan_id FOREIGN KEY (plan_id) REFERENCES WorkoutPlans(plan_id) ON DELETE SET NULL; -- If plan deleted, keep session log but remove plan link
ALTER TABLE UserWorkoutSessionLogs ADD CONSTRAINT fk_uwsl_session_id FOREIGN KEY (session_id) REFERENCES UserWorkoutSessions(session_id) ON DELETE CASCADE; -- If session deleted, delete logs for that session
ALTER TABLE UserWorkoutSessionLogs ADD CONSTRAINT fk_uwsl_exercise_id FOREIGN KEY (exercise_id) REFERENCES Exercises(exercise_id) ON DELETE RESTRICT; -- Prevent deleting exercise if it's in a log
ALTER TABLE UserWorkoutSessionLogs ADD CONSTRAINT fk_uwsl_plan_exercise_id FOREIGN KEY (plan_exercise_id) REFERENCES WorkoutPlanExercises(plan_exercise_id) ON DELETE SET NULL; -- If plan exercise detail deleted, keep log but remove link

-- Nutrition Module FKs
ALTER TABLE ServingUnits ADD CONSTRAINT fk_serving_units_food_item_id FOREIGN KEY (food_item_id) REFERENCES FoodItems(food_item_id) ON DELETE CASCADE; -- If food item deleted, delete its specific serving units
ALTER TABLE FoodItems ADD CONSTRAINT fk_food_items_created_by_user_id FOREIGN KEY (created_by_user_id) REFERENCES Users(user_id) ON DELETE SET NULL; -- If user deleted, keep food item but remove creator link (for public user items)
ALTER TABLE UserMeals ADD CONSTRAINT fk_user_meals_user_id FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE; -- If user deleted, delete their meal logs
ALTER TABLE MealFoodItems ADD CONSTRAINT fk_mfi_meal_id FOREIGN KEY (meal_id) REFERENCES UserMeals(meal_id) ON DELETE CASCADE; -- If meal deleted, delete its food items
ALTER TABLE MealFoodItems ADD CONSTRAINT fk_mfi_food_item_id FOREIGN KEY (food_item_id) REFERENCES FoodItems(food_item_id) ON DELETE RESTRICT; -- Prevent deleting food item if it's in a meal log
ALTER TABLE UserNutritionGoals ADD CONSTRAINT fk_ung_user_id FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE; -- If user deleted, delete their goals
ALTER TABLE UserBiometrics ADD CONSTRAINT fk_ub_user_id FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE; -- If user deleted, delete their biometric logs

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


-- Create base user

-- Commit the transaction
COMMIT;