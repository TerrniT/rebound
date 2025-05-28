package v1

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type mealRoutes struct {
	uc entity.MealUseCase
	l  logger.Interface
}

func NewMealRoutes(router fiber.Router, uc entity.MealUseCase, l logger.Interface) {
	r := &mealRoutes{
		uc: uc,
		l:  l,
	}

	group := router.Group("/meals")
	{
		group.Post("/", r.create)
		group.Get("/:id", r.getByID)
		group.Put("/:id", r.update)
		group.Delete("/:id", r.delete)
		group.Get("/", r.list)
		group.Get("/date/:date", r.getByDate)
	}
}

// @Summary Create a new meal
// @Description Create a new meal with the provided information
// @Tags meals
// @Accept json
// @Produce json
// @Param meal body entity.Meal true "Meal object"
// @Success 201 {object} entity.Meal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /meals [post]
func (r *mealRoutes) create(c *fiber.Ctx) error {
	meal := new(entity.Meal)
	if err := c.BodyParser(meal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := r.uc.CreateMeal(meal); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create meal",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(meal)
}

// @Summary Get meal by ID
// @Description Get meal information by its ID
// @Tags meals
// @Accept json
// @Produce json
// @Param id path int true "Meal ID"
// @Success 200 {object} entity.Meal
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /meals/{id} [get]
func (r *mealRoutes) getByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid meal ID",
		})
	}

	meal, err := r.uc.GetMeal(int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Meal not found",
		})
	}

	return c.JSON(meal)
}

// @Summary Update meal
// @Description Update meal information by its ID
// @Tags meals
// @Accept json
// @Produce json
// @Param id path int true "Meal ID"
// @Param meal body entity.Meal true "Meal object"
// @Success 200 {object} entity.Meal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /meals/{id} [put]
func (r *mealRoutes) update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid meal ID",
		})
	}

	meal := new(entity.Meal)
	if err := c.BodyParser(meal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	meal.ID = int64(id)
	if err := r.uc.UpdateMeal(meal); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update meal",
		})
	}

	return c.JSON(meal)
}

// @Summary Delete meal
// @Description Delete a meal by its ID
// @Tags meals
// @Accept json
// @Produce json
// @Param id path int true "Meal ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /meals/{id} [delete]
func (r *mealRoutes) delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid meal ID",
		})
	}

	if err := r.uc.DeleteMeal(int64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete meal",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary List meals
// @Description Get a list of meals with pagination
// @Tags meals
// @Accept json
// @Produce json
// @Param user_id query int true "User ID"
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Success 200 {array} entity.Meal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /meals [get]
func (r *mealRoutes) list(c *fiber.Ctx) error {
	userID := c.QueryInt("user_id")
	if userID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	startDate := time.Now().AddDate(0, -1, 0) // Default to last month
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		var err error
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start date format",
			})
		}
	}

	endDate := time.Now() // Default to now
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		var err error
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end date format",
			})
		}
	}

	meals, err := r.uc.ListMeals(int64(userID), startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list meals",
		})
	}

	return c.JSON(meals)
}

// @Summary Get meals by date
// @Description Get all meals for a specific date
// @Tags meals
// @Accept json
// @Produce json
// @Param user_id query int true "User ID"
// @Param date path string true "Date (YYYY-MM-DD)"
// @Success 200 {array} entity.Meal
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /meals/date/{date} [get]
func (r *mealRoutes) getByDate(c *fiber.Ctx) error {
	userID := c.QueryInt("user_id")
	if userID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	dateStr := c.Params("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format",
		})
	}

	meals, err := r.uc.GetDailyMeals(int64(userID), date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get meals for date",
		})
	}

	return c.JSON(meals)
}
