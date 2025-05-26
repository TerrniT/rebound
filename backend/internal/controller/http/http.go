// Package v1 implements routing paths. Each services in own file.
package http

// https://github.com/evrone/go-clean-template/tree/master/internal/controller/http/v1
import (
	"net/http"

	v1 "github.com/terrnit/rebound/backend/internal/controller/http/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "github.com/terrnit/rebound/backend/docs" // Swagger docs.
	"github.com/terrnit/rebound/backend/internal/entity"
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
	userUC entity.UserUseCase,
	nutritionUC entity.NutritionUseCase,
	mealUC entity.MealUseCase,
	productUC entity.ProductUseCase,
	authUC entity.AuthUseCase,
	l logger.Interface,
) *Router {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

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
		v1.NewNutritionRoutes(api, nutritionUC, l)
		v1.NewMealRoutes(api, mealUC, l)
		v1.NewProductRoutes(api, productUC, l)
		v1.NewAuthRoutes(api, authUC, l)
	}

	return &Router{
		app: app,
	}
}

func (r *Router) Run(addr string) error {
	return r.app.Listen(addr)
}

func (r *Router) Shutdown() error {
	return r.app.Shutdown()
}
