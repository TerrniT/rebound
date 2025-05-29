// Package v1 implements routing paths. Each services in own file.
package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "github.com/terrnit/rebound/backend/docs" // Swagger docs.
	v1 "github.com/terrnit/rebound/backend/internal/controller/router/v1"
	"github.com/terrnit/rebound/backend/internal/usecase"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type Router struct {
	app *fiber.App
}

// @title Rebound API
// @version 1.0
// @description Fitness application API for tracking nutrition, meals, and products
// @host localhost:8080
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func NewRouter(
	app *fiber.App,
	userUC *usecase.UserUseCase,
	foodItemUC *usecase.FoodItemUseCase,
	// mealUC *usecase.MealUseCase,
	// workoutPlanUC *usecase.WorkoutPlanUseCase,
	workoutSessionUC *usecase.WorkoutSessionUseCase,
	nutritionUC *usecase.NutritionUseCase,
	l logger.Interface,
) *Router {
	app.Use(cors.New())
	app.Use(fiberlogger.New(fiberlogger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	api := app.Group("/api")
	{
		v1.NewUserRoutes(api, userUC, l)
		v1.NewFoodItemRoutes(api, foodItemUC, l)
		// v1.NewMealRoutes(api, mealUC, l)
		v1.NewWorkoutSessionRoutes(api, workoutSessionUC, l)
		v1.NewNutritionRoutes(api, nutritionUC, l)
		// v1.NewExerciseRoutes()
	}

	return &Router{
		app: app,
	}
}
